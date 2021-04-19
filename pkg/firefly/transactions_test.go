package firefly_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/randaalex/finance_bot/pkg/firefly"
)

func TestFirefly_TransactionService_CreateTransaction(t *testing.T) {
	const (
		baseURL = "http://firefly/api/"
	)

	t.Run("valid request", func(t *testing.T) {
		req := firefly.CreateTransactionReq{
			Transactions: []firefly.CreateSubTransactionReq{
				{
					Type:          "withdrawal",
					Date:          "2021-05-03",
					Amount:        123.45,
					Description:   "test",
					SourceID:      6,
					DestinationID: 7,
					Notes:         "notes",
					ExternalID:    "external",
				},
			},
		}

		httpClient := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, "http://firefly/api/transactions", req.URL.String())

			expectedBody := `{"transactions":[{"type":"withdrawal","date":"2021-05-03","amount":123.45,"description":"test","source_id":6,"destination_id":7,"notes":"notes","external_id":"external"}]}` + "\n"

			assert.Equal(t, expectedBody, ReadCloserToString(req.Body))

			responseBody := `{
   "data":{
      "type":"transactions",
      "id":"42",
      "attributes":{
         "created_at":"2021-05-03T16:12:44+00:00",
         "updated_at":"2021-05-03T16:12:44+00:00",
         "user":1,
         "group_title":null,
         "transactions":[
            {
               "user":1,
               "transaction_journal_id":43,
               "type":"withdrawal",
               "date":"2021-05-03T00:00:00+00:00",
               "order":0,
               "currency_id":51,
               "currency_code":"BYN",
               "currency_name":"BYN",
               "currency_symbol":"BYN",
               "currency_decimal_places":2,
               "foreign_currency_id":0,
               "foreign_currency_code":null,
               "foreign_currency_symbol":null,
               "foreign_currency_decimal_places":0,
               "amount":"123.000000000000",
               "foreign_amount":null,
               "description":"test",
               "source_id":7,
               "source_name":"AlfaBank N1 (BYN)",
               "source_iban":"",
               "source_type":"Asset account",
               "destination_id":4,
               "destination_name":"Cash account",
               "destination_iban":null,
               "destination_type":"Cash account",
               "budget_id":0,
               "budget_name":null,
               "category_id":0,
               "category_name":null,
               "bill_id":0,
               "bill_name":null,
               "reconciled":false,
               "notes":null,
               "tags":[],
               "internal_reference":null,
               "external_id":null,
               "original_source":"ff3-v5.4.6|api-v1.4.0",
               "recurrence_id":null,
               "recurrence_total":null,
               "recurrence_count":null,
               "bunq_payment_id":null,
               "external_uri":null,
               "import_hash_v2":"1a7fe543e4c05b938ed678f193356837839b1f1bf43eb89e666360849adcfb84",
               "sepa_cc":null,
               "sepa_ct_op":null,
               "sepa_ct_id":null,
               "sepa_db":null,
               "sepa_country":null,
               "sepa_ep":null,
               "sepa_ci":null,
               "sepa_batch_id":null,
               "interest_date":null,
               "book_date":null,
               "process_date":null,
               "due_date":null,
               "payment_date":null,
               "invoice_date":null
            }
         ]
      },
      "links":{
         "self":"http:\/\/127.0.0.1:8081\/api\/v1\/transactions\/42",
         "0":{
            "rel":"self",
            "uri":"\/transactions\/42"
         }
      }
   }
}`

			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(responseBody)),
				Header:     make(http.Header),
			}
		})

		fireflyClient := firefly.NewClient(httpClient, baseURL)

		response, err := fireflyClient.Transactions.CreateTransaction(context.TODO(), req)

		expectedResponse := &firefly.Transaction{
			Data: firefly.TransactionData{
				Type: "transactions",
				ID:   "42",
				Attributes: firefly.TransactionAttributes{
					CreatedAt: "2021-05-03T16:12:44+00:00",
					Attributes: []firefly.TransactionSubTransaction{
						{
							User:                 1,
							TransactionJournalID: 43,
							Type:                 "withdrawal",
							Date:                 "2021-05-03T00:00:00+00:00",
							Amount:               "123.000000000000",
							Description:          "test",
							Order:                0,
						},
					},
				},
			},
		}

		assert.Equal(t, expectedResponse, response)
		assert.NoError(t, err)
	})

	t.Run("invalid request", func(t *testing.T) {
		req := firefly.CreateTransactionReq{
			Transactions: []firefly.CreateSubTransactionReq{},
		}

		httpClient := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, "http://firefly/api/transactions", req.URL.String())

			expectedBody := `{"transactions":[]}` + "\n"

			assert.Equal(t, expectedBody, ReadCloserToString(req.Body))

			return &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewBufferString("error")),
				Header:     make(http.Header),
			}
		})

		fireflyClient := firefly.NewClient(httpClient, baseURL)

		response, err := fireflyClient.Transactions.CreateTransaction(context.TODO(), req)

		assert.Nil(t, response)
		assert.Error(t, err, "invalid status code")
	})
}
