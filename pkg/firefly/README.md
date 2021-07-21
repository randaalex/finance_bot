# Go API client for firefly

This is the official documentation of the Firefly III API. You can find accompanying documentation on the website of Firefly III itself (see below). Please report any bugs or issues. This version of the API is live from version v4.7.9 and onwards. You may use the \"Authorize\" button to try the API below.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.4.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://firefly-iii.org](https://firefly-iii.org)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./firefly"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://demo.firefly-iii.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AboutApi* | [**GetAbout**](docs/AboutApi.md#getabout) | **Get** /api/v1/about | System information end point.
*AboutApi* | [**GetCurrentUser**](docs/AboutApi.md#getcurrentuser) | **Get** /api/v1/about/user | Currently authenticated user endpoint.
*AccountsApi* | [**DeleteAccount**](docs/AccountsApi.md#deleteaccount) | **Delete** /api/v1/accounts/{id} | Permanently delete account.
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /api/v1/accounts/{id} | Get single account.
*AccountsApi* | [**ListAccount**](docs/AccountsApi.md#listaccount) | **Get** /api/v1/accounts | List all accounts.
*AccountsApi* | [**ListAttachmentByAccount**](docs/AccountsApi.md#listattachmentbyaccount) | **Get** /api/v1/accounts/{id}/attachments | Lists all attachments.
*AccountsApi* | [**ListPiggyBankByAccount**](docs/AccountsApi.md#listpiggybankbyaccount) | **Get** /api/v1/accounts/{id}/piggy_banks | List all piggy banks related to the account.
*AccountsApi* | [**ListTransactionByAccount**](docs/AccountsApi.md#listtransactionbyaccount) | **Get** /api/v1/accounts/{id}/transactions | List all transactions related to the account.
*AccountsApi* | [**StoreAccount**](docs/AccountsApi.md#storeaccount) | **Post** /api/v1/accounts | Create new account.
*AccountsApi* | [**UpdateAccount**](docs/AccountsApi.md#updateaccount) | **Put** /api/v1/accounts/{id} | Update existing account.
*AttachmentsApi* | [**DeleteAttachment**](docs/AttachmentsApi.md#deleteattachment) | **Delete** /api/v1/attachments/{id} | Delete an attachment.
*AttachmentsApi* | [**DownloadAttachment**](docs/AttachmentsApi.md#downloadattachment) | **Get** /api/v1/attachments/{id}/download | Download a single attachment.
*AttachmentsApi* | [**GetAttachment**](docs/AttachmentsApi.md#getattachment) | **Get** /api/v1/attachments/{id} | Get a single attachment.
*AttachmentsApi* | [**ListAttachment**](docs/AttachmentsApi.md#listattachment) | **Get** /api/v1/attachments | List all attachments.
*AttachmentsApi* | [**StoreAttachment**](docs/AttachmentsApi.md#storeattachment) | **Post** /api/v1/attachments | Store a new attachment.
*AttachmentsApi* | [**UpdateAttachment**](docs/AttachmentsApi.md#updateattachment) | **Put** /api/v1/attachments/{id} | Update existing attachment.
*AttachmentsApi* | [**UploadAttachment**](docs/AttachmentsApi.md#uploadattachment) | **Post** /api/v1/attachments/{id}/upload | Upload an attachment.
*AutocompleteApi* | [**GetAccountsAC**](docs/AutocompleteApi.md#getaccountsac) | **Get** /api/v1/autocomplete/accounts | All accounts of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetBillsAC**](docs/AutocompleteApi.md#getbillsac) | **Get** /api/v1/autocomplete/bills | All bills of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetBudgetsAC**](docs/AutocompleteApi.md#getbudgetsac) | **Get** /api/v1/autocomplete/budgets | All budgets of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetCategoriesAC**](docs/AutocompleteApi.md#getcategoriesac) | **Get** /api/v1/autocomplete/categories | All categories of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetCurrenciesAC**](docs/AutocompleteApi.md#getcurrenciesac) | **Get** /api/v1/autocomplete/currencies | All currencies of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetCurrenciesCodeAC**](docs/AutocompleteApi.md#getcurrenciescodeac) | **Get** /api/v1/autocomplete/currencies-with-code | All currencies of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetObjectGroupsAC**](docs/AutocompleteApi.md#getobjectgroupsac) | **Get** /api/v1/autocomplete/object-groups | All object groups of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetPiggiesAC**](docs/AutocompleteApi.md#getpiggiesac) | **Get** /api/v1/autocomplete/piggy-banks | All piggy banks of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetPiggiesBalanceAC**](docs/AutocompleteApi.md#getpiggiesbalanceac) | **Get** /api/v1/autocomplete/piggy-banks-with-balance | All piggy banks of the user returned in a basic auto-complete array complemented with balance information.
*AutocompleteApi* | [**GetRuleGroupsAC**](docs/AutocompleteApi.md#getrulegroupsac) | **Get** /api/v1/autocomplete/rule-groups | All rule groups of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetRulesAC**](docs/AutocompleteApi.md#getrulesac) | **Get** /api/v1/autocomplete/rules | All rules of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetTagAC**](docs/AutocompleteApi.md#gettagac) | **Get** /api/v1/autocomplete/tags | All tags of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetTransactionTypesAC**](docs/AutocompleteApi.md#gettransactiontypesac) | **Get** /api/v1/autocomplete/transaction-types | All transaction types returned in a basic auto-complete array. English only.
*AutocompleteApi* | [**GetTransactionsAC**](docs/AutocompleteApi.md#gettransactionsac) | **Get** /api/v1/autocomplete/transactions | All transaction descriptions of the user returned in a basic auto-complete array.
*AutocompleteApi* | [**GetTransactionsIDAC**](docs/AutocompleteApi.md#gettransactionsidac) | **Get** /api/v1/autocomplete/transactions-with-id | All transactions, complemented with their ID, of the user returned in a basic auto-complete array.
*AvailableBudgetsApi* | [**DeleteAvailableBudget**](docs/AvailableBudgetsApi.md#deleteavailablebudget) | **Delete** /api/v1/available_budgets/{id} | Delete an available budget.
*AvailableBudgetsApi* | [**GetAvailableBudget**](docs/AvailableBudgetsApi.md#getavailablebudget) | **Get** /api/v1/available_budgets/{id} | Get a single available budget.
*AvailableBudgetsApi* | [**ListAvailableBudget**](docs/AvailableBudgetsApi.md#listavailablebudget) | **Get** /api/v1/available_budgets | List all available budget amounts.
*AvailableBudgetsApi* | [**StoreAvailableBudget**](docs/AvailableBudgetsApi.md#storeavailablebudget) | **Post** /api/v1/available_budgets | Store a new available budget
*AvailableBudgetsApi* | [**UpdateAvailableBudget**](docs/AvailableBudgetsApi.md#updateavailablebudget) | **Put** /api/v1/available_budgets/{id} | Update existing available budget, to change for example the date range of the amount or the amount itself.
*BillsApi* | [**DeleteBill**](docs/BillsApi.md#deletebill) | **Delete** /api/v1/bills/{id} | Delete a bill.
*BillsApi* | [**GetBill**](docs/BillsApi.md#getbill) | **Get** /api/v1/bills/{id} | Get a single bill.
*BillsApi* | [**ListAttachmentByBill**](docs/BillsApi.md#listattachmentbybill) | **Get** /api/v1/bills/{id}/attachments | List all attachments uploaded to the bill.
*BillsApi* | [**ListBill**](docs/BillsApi.md#listbill) | **Get** /api/v1/bills | List all bills.
*BillsApi* | [**ListRuleByBill**](docs/BillsApi.md#listrulebybill) | **Get** /api/v1/bills/{id}/rules | List all rules associated with the bill.
*BillsApi* | [**ListTransactionByBill**](docs/BillsApi.md#listtransactionbybill) | **Get** /api/v1/bills/{id}/transactions | List all transactions associated with the  bill.
*BillsApi* | [**StoreBill**](docs/BillsApi.md#storebill) | **Post** /api/v1/bills | Store a new bill
*BillsApi* | [**UpdateBill**](docs/BillsApi.md#updatebill) | **Put** /api/v1/bills/{id} | Update existing bill.
*BudgetsApi* | [**DeleteBudget**](docs/BudgetsApi.md#deletebudget) | **Delete** /api/v1/budgets/{id} | Delete a budget.
*BudgetsApi* | [**DeleteBudgetLimit**](docs/BudgetsApi.md#deletebudgetlimit) | **Delete** /api/v1/budgets/limits/{id} | Delete a budget limit.
*BudgetsApi* | [**GetBudget**](docs/BudgetsApi.md#getbudget) | **Get** /api/v1/budgets/{id} | Get a single budget.
*BudgetsApi* | [**GetBudgetLimit**](docs/BudgetsApi.md#getbudgetlimit) | **Get** /api/v1/budgets/limits/{id} | Get single budget limit.
*BudgetsApi* | [**ListAttachmentByBudget**](docs/BudgetsApi.md#listattachmentbybudget) | **Get** /api/v1/budgets/{id}/attachments | Lists all attachments.
*BudgetsApi* | [**ListBudget**](docs/BudgetsApi.md#listbudget) | **Get** /api/v1/budgets | List all budgets.
*BudgetsApi* | [**ListBudgetLimitByBudget**](docs/BudgetsApi.md#listbudgetlimitbybudget) | **Get** /api/v1/budgets/{id}/limits | Get all limits
*BudgetsApi* | [**ListTransactionByBudget**](docs/BudgetsApi.md#listtransactionbybudget) | **Get** /api/v1/budgets/{id}/transactions | All transactions to a budget.
*BudgetsApi* | [**ListTransactionByBudgetLimit**](docs/BudgetsApi.md#listtransactionbybudgetlimit) | **Get** /api/v1/budgets/limits/{id}/transactions | List all transactions by a budget limit ID.
*BudgetsApi* | [**StoreBudget**](docs/BudgetsApi.md#storebudget) | **Post** /api/v1/budgets | Store a new budget
*BudgetsApi* | [**StoreBudgetLimit**](docs/BudgetsApi.md#storebudgetlimit) | **Post** /api/v1/budgets/{id}/limits | Store new budget limit.
*BudgetsApi* | [**UpdateBudget**](docs/BudgetsApi.md#updatebudget) | **Put** /api/v1/budgets/{id} | Update existing budget.
*BudgetsApi* | [**UpdateBudgetLimit**](docs/BudgetsApi.md#updatebudgetlimit) | **Put** /api/v1/budgets/limits/{id} | Update existing budget limit.
*CategoriesApi* | [**DeleteCategory**](docs/CategoriesApi.md#deletecategory) | **Delete** /api/v1/categories/{id} | Delete a category.
*CategoriesApi* | [**GetCategory**](docs/CategoriesApi.md#getcategory) | **Get** /api/v1/categories/{id} | Get a single category.
*CategoriesApi* | [**ListAttachmentByCategory**](docs/CategoriesApi.md#listattachmentbycategory) | **Get** /api/v1/categories/{id}/attachments | Lists all attachments.
*CategoriesApi* | [**ListCategory**](docs/CategoriesApi.md#listcategory) | **Get** /api/v1/categories | List all categories.
*CategoriesApi* | [**ListTransactionByCategory**](docs/CategoriesApi.md#listtransactionbycategory) | **Get** /api/v1/categories/{id}/transactions | List all transactions in a category.
*CategoriesApi* | [**StoreCategory**](docs/CategoriesApi.md#storecategory) | **Post** /api/v1/categories | Store a new category
*CategoriesApi* | [**UpdateCategory**](docs/CategoriesApi.md#updatecategory) | **Put** /api/v1/categories/{id} | Update existing category.
*ChartsApi* | [**GetChartABOverview**](docs/ChartsApi.md#getchartaboverview) | **Get** /api/v1/chart/ab/overview/{id} | Dashboard chart with an overview of the available budget.
*ChartsApi* | [**GetChartAccountExpense**](docs/ChartsApi.md#getchartaccountexpense) | **Get** /api/v1/chart/account/expense | Dashboard chart with expense account balance information.
*ChartsApi* | [**GetChartAccountOverview**](docs/ChartsApi.md#getchartaccountoverview) | **Get** /api/v1/chart/account/overview | Dashboard chart with asset account balance information.
*ChartsApi* | [**GetChartAccountRevenue**](docs/ChartsApi.md#getchartaccountrevenue) | **Get** /api/v1/chart/account/revenue | Dashboard chart with revenue account balance information.
*ChartsApi* | [**GetChartCategoryOverview**](docs/ChartsApi.md#getchartcategoryoverview) | **Get** /api/v1/chart/category/overview | Dashboard chart with an overview of the users categories.
*ConfigurationApi* | [**GetConfiguration**](docs/ConfigurationApi.md#getconfiguration) | **Get** /api/v1/configuration | Get Firefly III system configuration.
*ConfigurationApi* | [**SetConfiguration**](docs/ConfigurationApi.md#setconfiguration) | **Post** /api/v1/configuration/{name} | Update configuration
*CurrenciesApi* | [**DefaultCurrency**](docs/CurrenciesApi.md#defaultcurrency) | **Post** /api/v1/currencies/{code}/default | Make currency default currency.
*CurrenciesApi* | [**DeleteCurrency**](docs/CurrenciesApi.md#deletecurrency) | **Delete** /api/v1/currencies/{code} | Delete a currency.
*CurrenciesApi* | [**DisableCurrency**](docs/CurrenciesApi.md#disablecurrency) | **Post** /api/v1/currencies/{code}/disable | Disable a currency.
*CurrenciesApi* | [**EnableCurrency**](docs/CurrenciesApi.md#enablecurrency) | **Post** /api/v1/currencies/{code}/enable | Enable a single currency.
*CurrenciesApi* | [**GetCurrency**](docs/CurrenciesApi.md#getcurrency) | **Get** /api/v1/currencies/{code} | Get a single currency.
*CurrenciesApi* | [**GetDefaultCurrency**](docs/CurrenciesApi.md#getdefaultcurrency) | **Get** /api/v1/currencies/default | Get the user&#39;s default currency.
*CurrenciesApi* | [**ListAccountByCurrency**](docs/CurrenciesApi.md#listaccountbycurrency) | **Get** /api/v1/currencies/{code}/accounts | List all accounts with this currency.
*CurrenciesApi* | [**ListAvailableBudgetByCurrency**](docs/CurrenciesApi.md#listavailablebudgetbycurrency) | **Get** /api/v1/currencies/{code}/available_budgets | List all available budgets with this currency.
*CurrenciesApi* | [**ListBillByCurrency**](docs/CurrenciesApi.md#listbillbycurrency) | **Get** /api/v1/currencies/{code}/bills | List all bills with this currency.
*CurrenciesApi* | [**ListBudgetLimitByCurrency**](docs/CurrenciesApi.md#listbudgetlimitbycurrency) | **Get** /api/v1/currencies/{code}/budget_limits | List all budget limits with this currency
*CurrenciesApi* | [**ListCurrency**](docs/CurrenciesApi.md#listcurrency) | **Get** /api/v1/currencies | List all currencies.
*CurrenciesApi* | [**ListExchangeRateByCurrency**](docs/CurrenciesApi.md#listexchangeratebycurrency) | **Get** /api/v1/currencies/{code}/cer | List all known exchange rates with (from or to) this currency.
*CurrenciesApi* | [**ListRecurrenceByCurrency**](docs/CurrenciesApi.md#listrecurrencebycurrency) | **Get** /api/v1/currencies/{code}/recurrences | List all recurring transactions with this currency.
*CurrenciesApi* | [**ListRuleByCurrency**](docs/CurrenciesApi.md#listrulebycurrency) | **Get** /api/v1/currencies/{code}/rules | List all rules with this currency.
*CurrenciesApi* | [**ListTransactionByCurrency**](docs/CurrenciesApi.md#listtransactionbycurrency) | **Get** /api/v1/currencies/{code}/transactions | List all transactions with this currency.
*CurrenciesApi* | [**StoreCurrency**](docs/CurrenciesApi.md#storecurrency) | **Post** /api/v1/currencies | Store a new currency
*CurrenciesApi* | [**UpdateCurrency**](docs/CurrenciesApi.md#updatecurrency) | **Put** /api/v1/currencies/{code} | Update existing currency.
*DataApi* | [**DestroyData**](docs/DataApi.md#destroydata) | **Delete** /api/v1/data/destroy | Endpoint to destroy user data
*ImportApi* | [**GetImport**](docs/ImportApi.md#getimport) | **Get** /api/v1/import/{key} | Show info on a single import
*ImportApi* | [**ListImport**](docs/ImportApi.md#listimport) | **Get** /api/v1/import/list | List al imports
*ImportApi* | [**ListTransactionByImport**](docs/ImportApi.md#listtransactionbyimport) | **Get** /api/v1/import/{key}/transactions | List all transactions related to the import job. The correlation is made through the tag.
*LinksApi* | [**DeleteLinkType**](docs/LinksApi.md#deletelinktype) | **Delete** /api/v1/link_types/{id} | Permanently delete link type.
*LinksApi* | [**DeleteTransactionLink**](docs/LinksApi.md#deletetransactionlink) | **Delete** /api/v1/transaction_links/{id} | Permanently delete link between transactions.
*LinksApi* | [**GetLinkType**](docs/LinksApi.md#getlinktype) | **Get** /api/v1/link_types/{id} | Get single a link type.
*LinksApi* | [**GetTransactionLink**](docs/LinksApi.md#gettransactionlink) | **Get** /api/v1/transaction_links/{id} | Get a single link.
*LinksApi* | [**ListLinkType**](docs/LinksApi.md#listlinktype) | **Get** /api/v1/link_types | List all types of links.
*LinksApi* | [**ListTransactionByLinkType**](docs/LinksApi.md#listtransactionbylinktype) | **Get** /api/v1/link_types/{id}/transactions | List all transactions under this link type.
*LinksApi* | [**ListTransactionLink**](docs/LinksApi.md#listtransactionlink) | **Get** /api/v1/transaction_links | List all transaction links.
*LinksApi* | [**StoreLinkType**](docs/LinksApi.md#storelinktype) | **Post** /api/v1/link_types | Create a new link type
*LinksApi* | [**StoreTransactionLink**](docs/LinksApi.md#storetransactionlink) | **Post** /api/v1/transaction_links | Create a new link between transactions
*LinksApi* | [**UpdateLinkType**](docs/LinksApi.md#updatelinktype) | **Put** /api/v1/link_types/{id} | Update existing link type.
*LinksApi* | [**UpdateTransactionLink**](docs/LinksApi.md#updatetransactionlink) | **Put** /api/v1/transaction_links/{id} | Update an existing link between transactions.
*PiggyBanksApi* | [**DeletePiggyBank**](docs/PiggyBanksApi.md#deletepiggybank) | **Delete** /api/v1/piggy_banks/{id} | Delete a piggy bank.
*PiggyBanksApi* | [**GetPiggyBank**](docs/PiggyBanksApi.md#getpiggybank) | **Get** /api/v1/piggy_banks/{id} | Get a single piggy bank.
*PiggyBanksApi* | [**ListAttachmentByPiggyBank**](docs/PiggyBanksApi.md#listattachmentbypiggybank) | **Get** /api/v1/piggy_banks/{id}/attachments | Lists all attachments.
*PiggyBanksApi* | [**ListEventByPiggyBank**](docs/PiggyBanksApi.md#listeventbypiggybank) | **Get** /api/v1/piggy_banks/{id}/events | List all events linked to a piggy bank.
*PiggyBanksApi* | [**ListPiggyBank**](docs/PiggyBanksApi.md#listpiggybank) | **Get** /api/v1/piggy_banks | List all piggy banks.
*PiggyBanksApi* | [**StorePiggyBank**](docs/PiggyBanksApi.md#storepiggybank) | **Post** /api/v1/piggy_banks | Store a new piggy bank
*PiggyBanksApi* | [**UpdatePiggyBank**](docs/PiggyBanksApi.md#updatepiggybank) | **Put** /api/v1/piggy_banks/{id} | Update existing piggy bank.
*PreferencesApi* | [**GetPreference**](docs/PreferencesApi.md#getpreference) | **Get** /api/v1/preferences/{name} | Return a single preference.
*PreferencesApi* | [**ListPreference**](docs/PreferencesApi.md#listpreference) | **Get** /api/v1/preferences | List all users preferences.
*PreferencesApi* | [**UpdatePreference**](docs/PreferencesApi.md#updatepreference) | **Put** /api/v1/preferences/{name} | Update preference
*RecurrencesApi* | [**DeleteRecurrence**](docs/RecurrencesApi.md#deleterecurrence) | **Delete** /api/v1/recurrences/{id} | Delete a recurring transaction.
*RecurrencesApi* | [**GetRecurrence**](docs/RecurrencesApi.md#getrecurrence) | **Get** /api/v1/recurrences/{id} | Get a single recurring transaction.
*RecurrencesApi* | [**ListRecurrence**](docs/RecurrencesApi.md#listrecurrence) | **Get** /api/v1/recurrences | List all recurring transactions.
*RecurrencesApi* | [**ListTransactionByRecurrence**](docs/RecurrencesApi.md#listtransactionbyrecurrence) | **Get** /api/v1/recurrences/{id}/transactions | List all transactions created by a recurring transaction.
*RecurrencesApi* | [**StoreRecurrence**](docs/RecurrencesApi.md#storerecurrence) | **Post** /api/v1/recurrences | Store a new recurring transaction
*RecurrencesApi* | [**TriggerRecurrence**](docs/RecurrencesApi.md#triggerrecurrence) | **Post** /api/v1/recurrences/trigger | Trigger the creation of recurring transactions (like a cron job).
*RecurrencesApi* | [**UpdateRecurrence**](docs/RecurrencesApi.md#updaterecurrence) | **Put** /api/v1/recurrences/{id} | Update existing recurring transaction.
*RuleGroupsApi* | [**DeleteRuleGroup**](docs/RuleGroupsApi.md#deleterulegroup) | **Delete** /api/v1/rule_groups/{id} | Delete a rule group.
*RuleGroupsApi* | [**FireRuleGroup**](docs/RuleGroupsApi.md#firerulegroup) | **Post** /api/v1/rule_groups/{id}/trigger | Fire the rule group on your transactions.
*RuleGroupsApi* | [**GetRuleGroup**](docs/RuleGroupsApi.md#getrulegroup) | **Get** /api/v1/rule_groups/{id} | Get a single rule group.
*RuleGroupsApi* | [**ListRuleByGroup**](docs/RuleGroupsApi.md#listrulebygroup) | **Get** /api/v1/rule_groups/{id}/rules | List rules in this rule group.
*RuleGroupsApi* | [**ListRuleGroup**](docs/RuleGroupsApi.md#listrulegroup) | **Get** /api/v1/rule_groups | List all rule groups.
*RuleGroupsApi* | [**StoreRuleGroup**](docs/RuleGroupsApi.md#storerulegroup) | **Post** /api/v1/rule_groups | Store a new rule group.
*RuleGroupsApi* | [**TestRuleGroup**](docs/RuleGroupsApi.md#testrulegroup) | **Get** /api/v1/rule_groups/{id}/test | Test which transactions would be hit by the rule group. No changes will be made.
*RuleGroupsApi* | [**UpdateRuleGroup**](docs/RuleGroupsApi.md#updaterulegroup) | **Put** /api/v1/rule_groups/{id} | Update existing rule group.
*RulesApi* | [**DeleteRule**](docs/RulesApi.md#deleterule) | **Delete** /api/v1/rules/{id} | Delete an rule.
*RulesApi* | [**FireRule**](docs/RulesApi.md#firerule) | **Post** /api/v1/rules/{id}/trigger | Fire the rule on your transactions.
*RulesApi* | [**GetRule**](docs/RulesApi.md#getrule) | **Get** /api/v1/rules/{id} | Get a single rule.
*RulesApi* | [**ListRule**](docs/RulesApi.md#listrule) | **Get** /api/v1/rules | List all rules.
*RulesApi* | [**StoreRule**](docs/RulesApi.md#storerule) | **Post** /api/v1/rules | Store a new rule
*RulesApi* | [**TestRule**](docs/RulesApi.md#testrule) | **Get** /api/v1/rules/{id}/test | Test which transactions would be hit by the rule. No changes will be made.
*RulesApi* | [**UpdateRule**](docs/RulesApi.md#updaterule) | **Put** /api/v1/rules/{id} | Update existing rule.
*SearchApi* | [**SearchAccounts**](docs/SearchApi.md#searchaccounts) | **Get** /api/v1/search/accounts | Search for accounts
*SearchApi* | [**SearchTransactions**](docs/SearchApi.md#searchtransactions) | **Get** /api/v1/search/transactions | Search for transactions
*SummaryApi* | [**GetBasicSummary**](docs/SummaryApi.md#getbasicsummary) | **Get** /api/v1/summary/basic | Returns basic sums of the users data.
*TagsApi* | [**DeleteTag**](docs/TagsApi.md#deletetag) | **Delete** /api/v1/tags/{tag} | Delete an tag.
*TagsApi* | [**GetTag**](docs/TagsApi.md#gettag) | **Get** /api/v1/tags/{tag} | Get a single tag.
*TagsApi* | [**GetTagCloud**](docs/TagsApi.md#gettagcloud) | **Get** /api/v1/tag-cloud | Returns a basic tag cloud.
*TagsApi* | [**ListAttachmentByTag**](docs/TagsApi.md#listattachmentbytag) | **Get** /api/v1/tags/{tag}/attachments | Lists all attachments.
*TagsApi* | [**ListTag**](docs/TagsApi.md#listtag) | **Get** /api/v1/tags | List all tags.
*TagsApi* | [**ListTransactionByTag**](docs/TagsApi.md#listtransactionbytag) | **Get** /api/v1/tags/{tag}/transactions | List all transactions with this tag.
*TagsApi* | [**StoreTag**](docs/TagsApi.md#storetag) | **Post** /api/v1/tags | Store a new tag
*TagsApi* | [**UpdateTag**](docs/TagsApi.md#updatetag) | **Put** /api/v1/tags/{tag} | Update existing tag.
*TransactionsApi* | [**DeleteTransaction**](docs/TransactionsApi.md#deletetransaction) | **Delete** /api/v1/transactions/{id} | Delete a transaction.
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /api/v1/transactions/{id} | Get a single transaction.
*TransactionsApi* | [**GetTransactionByJournal**](docs/TransactionsApi.md#gettransactionbyjournal) | **Get** /api/v1/transaction-journals/{id} | Get a single transaction, based on one of the underlying transaction journals.
*TransactionsApi* | [**ListAttachmentByTransaction**](docs/TransactionsApi.md#listattachmentbytransaction) | **Get** /api/v1/transactions/{id}/attachments | Lists all attachments.
*TransactionsApi* | [**ListEventByTransaction**](docs/TransactionsApi.md#listeventbytransaction) | **Get** /api/v1/transactions/{id}/piggy_bank_events | Lists all piggy bank events.
*TransactionsApi* | [**ListTransaction**](docs/TransactionsApi.md#listtransaction) | **Get** /api/v1/transactions | List all the user&#39;s transactions. 
*TransactionsApi* | [**StoreTransaction**](docs/TransactionsApi.md#storetransaction) | **Post** /api/v1/transactions | Store a new transaction
*TransactionsApi* | [**UpdateTransaction**](docs/TransactionsApi.md#updatetransaction) | **Put** /api/v1/transactions/{id} | Update existing transaction.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /api/v1/users/{id} | Delete a user.
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /api/v1/users/{id} | Get a single user.
*UsersApi* | [**ListUser**](docs/UsersApi.md#listuser) | **Get** /api/v1/users | List all users.
*UsersApi* | [**StoreUser**](docs/UsersApi.md#storeuser) | **Post** /api/v1/users | Store a new user
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /api/v1/users/{id} | Update an existing user&#39;s information.


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountArray](docs/AccountArray.md)
 - [AccountRead](docs/AccountRead.md)
 - [AccountSearchFieldFilter](docs/AccountSearchFieldFilter.md)
 - [AccountSingle](docs/AccountSingle.md)
 - [AccountTypeFilter](docs/AccountTypeFilter.md)
 - [AccountTypeProperty](docs/AccountTypeProperty.md)
 - [Attachment](docs/Attachment.md)
 - [AttachmentArray](docs/AttachmentArray.md)
 - [AttachmentRead](docs/AttachmentRead.md)
 - [AttachmentSingle](docs/AttachmentSingle.md)
 - [AutocompleteAccount](docs/AutocompleteAccount.md)
 - [AutocompleteBill](docs/AutocompleteBill.md)
 - [AutocompleteBudget](docs/AutocompleteBudget.md)
 - [AutocompleteCategory](docs/AutocompleteCategory.md)
 - [AutocompleteCurrency](docs/AutocompleteCurrency.md)
 - [AutocompleteCurrencyCode](docs/AutocompleteCurrencyCode.md)
 - [AutocompleteObjectGroup](docs/AutocompleteObjectGroup.md)
 - [AutocompletePiggy](docs/AutocompletePiggy.md)
 - [AutocompletePiggyBalance](docs/AutocompletePiggyBalance.md)
 - [AutocompleteRule](docs/AutocompleteRule.md)
 - [AutocompleteRuleGroup](docs/AutocompleteRuleGroup.md)
 - [AutocompleteTag](docs/AutocompleteTag.md)
 - [AutocompleteTransaction](docs/AutocompleteTransaction.md)
 - [AutocompleteTransactionID](docs/AutocompleteTransactionID.md)
 - [AutocompleteTransactionType](docs/AutocompleteTransactionType.md)
 - [AvailableBudget](docs/AvailableBudget.md)
 - [AvailableBudgetArray](docs/AvailableBudgetArray.md)
 - [AvailableBudgetRead](docs/AvailableBudgetRead.md)
 - [AvailableBudgetSingle](docs/AvailableBudgetSingle.md)
 - [BasicSummaryEntry](docs/BasicSummaryEntry.md)
 - [Bill](docs/Bill.md)
 - [BillArray](docs/BillArray.md)
 - [BillPaidDates](docs/BillPaidDates.md)
 - [BillRead](docs/BillRead.md)
 - [BillSingle](docs/BillSingle.md)
 - [Budget](docs/Budget.md)
 - [BudgetArray](docs/BudgetArray.md)
 - [BudgetLimit](docs/BudgetLimit.md)
 - [BudgetLimitArray](docs/BudgetLimitArray.md)
 - [BudgetLimitRead](docs/BudgetLimitRead.md)
 - [BudgetLimitSingle](docs/BudgetLimitSingle.md)
 - [BudgetRead](docs/BudgetRead.md)
 - [BudgetSingle](docs/BudgetSingle.md)
 - [BudgetSpent](docs/BudgetSpent.md)
 - [Category](docs/Category.md)
 - [CategoryArray](docs/CategoryArray.md)
 - [CategoryEarned](docs/CategoryEarned.md)
 - [CategoryRead](docs/CategoryRead.md)
 - [CategorySingle](docs/CategorySingle.md)
 - [CategorySpent](docs/CategorySpent.md)
 - [ChartDataPoint](docs/ChartDataPoint.md)
 - [ChartDataSet](docs/ChartDataSet.md)
 - [Configuration](docs/Configuration.md)
 - [ConfigurationData](docs/ConfigurationData.md)
 - [ConfigurationUpdate](docs/ConfigurationUpdate.md)
 - [Currency](docs/Currency.md)
 - [CurrencyArray](docs/CurrencyArray.md)
 - [CurrencyRead](docs/CurrencyRead.md)
 - [CurrencySingle](docs/CurrencySingle.md)
 - [DataDestroyObject](docs/DataDestroyObject.md)
 - [ExchangeRate](docs/ExchangeRate.md)
 - [ExchangeRateArray](docs/ExchangeRateArray.md)
 - [ExchangeRateAttributes](docs/ExchangeRateAttributes.md)
 - [ImportJob](docs/ImportJob.md)
 - [ImportJobArray](docs/ImportJobArray.md)
 - [ImportJobAttributes](docs/ImportJobAttributes.md)
 - [ImportJobSingle](docs/ImportJobSingle.md)
 - [LinkType](docs/LinkType.md)
 - [LinkTypeArray](docs/LinkTypeArray.md)
 - [LinkTypeRead](docs/LinkTypeRead.md)
 - [LinkTypeSingle](docs/LinkTypeSingle.md)
 - [Meta](docs/Meta.md)
 - [MetaPagination](docs/MetaPagination.md)
 - [ObjectLink](docs/ObjectLink.md)
 - [ObjectLink0](docs/ObjectLink0.md)
 - [PageLink](docs/PageLink.md)
 - [PiggyBank](docs/PiggyBank.md)
 - [PiggyBankArray](docs/PiggyBankArray.md)
 - [PiggyBankEvent](docs/PiggyBankEvent.md)
 - [PiggyBankEventArray](docs/PiggyBankEventArray.md)
 - [PiggyBankEventRead](docs/PiggyBankEventRead.md)
 - [PiggyBankRead](docs/PiggyBankRead.md)
 - [PiggyBankSingle](docs/PiggyBankSingle.md)
 - [Preference](docs/Preference.md)
 - [PreferenceArray](docs/PreferenceArray.md)
 - [PreferenceRead](docs/PreferenceRead.md)
 - [PreferenceSingle](docs/PreferenceSingle.md)
 - [Recurrence](docs/Recurrence.md)
 - [RecurrenceArray](docs/RecurrenceArray.md)
 - [RecurrenceRead](docs/RecurrenceRead.md)
 - [RecurrenceRepetition](docs/RecurrenceRepetition.md)
 - [RecurrenceSingle](docs/RecurrenceSingle.md)
 - [RecurrenceTransaction](docs/RecurrenceTransaction.md)
 - [Rule](docs/Rule.md)
 - [RuleAction](docs/RuleAction.md)
 - [RuleArray](docs/RuleArray.md)
 - [RuleGroup](docs/RuleGroup.md)
 - [RuleGroupArray](docs/RuleGroupArray.md)
 - [RuleGroupRead](docs/RuleGroupRead.md)
 - [RuleGroupSingle](docs/RuleGroupSingle.md)
 - [RuleRead](docs/RuleRead.md)
 - [RuleSingle](docs/RuleSingle.md)
 - [RuleTrigger](docs/RuleTrigger.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [SystemInfoData](docs/SystemInfoData.md)
 - [TagArray](docs/TagArray.md)
 - [TagCloud](docs/TagCloud.md)
 - [TagCloudTag](docs/TagCloudTag.md)
 - [TagModel](docs/TagModel.md)
 - [TagRead](docs/TagRead.md)
 - [TagSingle](docs/TagSingle.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionArray](docs/TransactionArray.md)
 - [TransactionLink](docs/TransactionLink.md)
 - [TransactionLinkArray](docs/TransactionLinkArray.md)
 - [TransactionLinkRead](docs/TransactionLinkRead.md)
 - [TransactionLinkSingle](docs/TransactionLinkSingle.md)
 - [TransactionRead](docs/TransactionRead.md)
 - [TransactionSingle](docs/TransactionSingle.md)
 - [TransactionSplit](docs/TransactionSplit.md)
 - [TransactionTypeFilter](docs/TransactionTypeFilter.md)
 - [User](docs/User.md)
 - [UserArray](docs/UserArray.md)
 - [UserRead](docs/UserRead.md)
 - [UserSingle](docs/UserSingle.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorErrors](docs/ValidationErrorErrors.md)


## Documentation For Authorization



### firefly_iii_auth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://demo.firefly-iii.org/oauth/authorize
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

james@firefly-iii.org
