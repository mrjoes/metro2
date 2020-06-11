// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

type field struct {
	Start    int
	Length   int
	Type     int
	Required string
}

type specification struct {
	Key   int
	Name  string
	Field field
}

const (
	nullable            = ""
	required            = "Y"
	applicable          = "A"
	timestampSizeStr    = "14"
	dateSizeStr         = "8"
	packedTimestampSize = 8
	packedDateSize      = 5
	int64size           = 8
)

// field types
const (
	alphanumeric = 1 << iota
	alpha
	numeric
	timestamp
	date
	descriptor
	packedTimestamp
	packedDate
	packedNumber
)

// field type options
const (
	zeroFill = 1 << 14
	omitted  = 1 << 15
)

var (
	baseSegmentCharacterFormat = map[string]field{
		"BlockDescriptorWord":           {0, 4, numeric | omitted, applicable},
		"RecordDescriptorWord":          {0, 4, numeric, required},
		"ProcessingIndicator":           {4, 1, numeric, nullable},
		"TimeStamp":                     {5, 14, timestamp, nullable},
		"Reserved1":                     {19, 1, alphanumeric | zeroFill, nullable},
		"IdentificationNumber":          {20, 20, alphanumeric, required},
		"CycleIdentifier":               {40, 2, alphanumeric, applicable},
		"ConsumerAccountNumber":         {42, 30, alphanumeric, required},
		"PortfolioType":                 {72, 1, alphanumeric, required},
		"AccountType":                   {73, 2, alphanumeric, required},
		"DateOpened":                    {75, 8, date, required},
		"CreditLimit":                   {83, 9, numeric | zeroFill, applicable},
		"HighestCredit":                 {92, 9, numeric, required},
		"TermsDuration":                 {101, 3, alphanumeric, required},
		"TermsFrequency":                {104, 1, alphanumeric, applicable},
		"ScheduledMonthlyPaymentAmount": {105, 9, numeric, applicable},
		"ActualPaymentAmount":           {114, 9, numeric, applicable},
		"AccountStatus":                 {123, 2, alphanumeric, required},
		"PaymentRating":                 {125, 1, alphanumeric, applicable},
		"PaymentHistoryProfile":         {126, 24, alphanumeric, required},
		"SpecialComment":                {150, 2, alphanumeric, applicable},
		"ComplianceConditionCode":       {152, 2, alphanumeric, applicable},
		"CurrentBalance":                {154, 9, numeric, required},
		"AmountPastDue":                 {163, 9, numeric, applicable},
		"OriginalChargeOffAmount":       {172, 9, numeric, applicable},
		"DateAccountInformation":        {181, 8, date, required},
		"DateFirstDelinquency":          {189, 8, date, applicable},
		"DateClosed":                    {197, 8, date | zeroFill, applicable},
		"DateLastPayment":               {205, 8, date, applicable},
		"InterestTypeIndicator":         {213, 1, alphanumeric, nullable},
		"Reserved2":                     {214, 17, alphanumeric, nullable},
		"Surname":                       {231, 25, alphanumeric, required},
		"FirstName":                     {256, 20, alphanumeric, required},
		"MiddleName":                    {276, 20, alphanumeric, applicable},
		"GenerationCode":                {296, 1, alphanumeric, applicable},
		"SocialSecurityNumber":          {297, 9, numeric, required},
		"DateBirth":                     {306, 8, date, required},
		"TelephoneNumber":               {314, 10, numeric, nullable},
		"ECOACode":                      {324, 1, alphanumeric, required},
		"ConsumerInformationIndicator":  {325, 2, alphanumeric, applicable},
		"CountryCode":                   {327, 2, alphanumeric, nullable},
		"FirstLineAddress":              {329, 32, alphanumeric, required},
		"SecondLineAddress":             {361, 32, alphanumeric, applicable},
		"City":                          {393, 20, alphanumeric, required},
		"State":                         {413, 2, alphanumeric, required},
		"ZipCode":                       {415, 9, alphanumeric, required},
		"AddressIndicator":              {424, 1, alphanumeric, nullable},
		"ResidenceCode":                 {425, 1, alphanumeric, nullable},
	}
	baseSegmentPackedFormat = map[string]field{
		"BlockDescriptorWord":           {0, 4, descriptor | omitted, applicable},
		"RecordDescriptorWord":          {0, 4, descriptor, required},
		"ProcessingIndicator":           {4, 1, numeric, nullable},
		"TimeStamp":                     {5, 8, packedTimestamp, nullable},
		"Reserved1":                     {13, 1, alphanumeric | zeroFill, nullable},
		"IdentificationNumber":          {14, 20, alphanumeric, required},
		"CycleIdentifier":               {34, 2, alphanumeric, applicable},
		"ConsumerAccountNumber":         {36, 30, alphanumeric, required},
		"PortfolioType":                 {66, 1, alphanumeric, required},
		"AccountType":                   {67, 2, alphanumeric, required},
		"DateOpened":                    {69, 5, packedDate, required},
		"CreditLimit":                   {74, 5, packedNumber | zeroFill, applicable},
		"HighestCredit":                 {79, 5, packedNumber, required},
		"TermsDuration":                 {84, 3, alphanumeric, required},
		"TermsFrequency":                {87, 1, alphanumeric, applicable},
		"ScheduledMonthlyPaymentAmount": {88, 5, packedNumber, applicable},
		"ActualPaymentAmount":           {93, 5, packedNumber, applicable},
		"AccountStatus":                 {98, 2, alphanumeric, required},
		"PaymentRating":                 {100, 1, alphanumeric, applicable},
		"PaymentHistoryProfile":         {101, 24, alphanumeric, required},
		"SpecialComment":                {125, 2, alphanumeric, applicable},
		"ComplianceConditionCode":       {127, 2, alphanumeric, applicable},
		"CurrentBalance":                {129, 5, packedNumber, required},
		"AmountPastDue":                 {134, 5, packedNumber, applicable},
		"OriginalChargeOffAmount":       {139, 5, packedNumber, applicable},
		"DateAccountInformation":        {144, 5, packedDate, required},
		"DateFirstDelinquency":          {149, 5, packedDate, applicable},
		"DateClosed":                    {154, 5, packedDate | zeroFill, applicable},
		"DateLastPayment":               {159, 5, packedDate, applicable},
		"InterestTypeIndicator":         {164, 1, alphanumeric, nullable},
		"Reserved2":                     {165, 17, alphanumeric, nullable},
		"Surname":                       {182, 25, alphanumeric, required},
		"FirstName":                     {207, 20, alphanumeric, required},
		"MiddleName":                    {227, 20, alphanumeric, applicable},
		"GenerationCode":                {247, 1, alphanumeric, applicable},
		"SocialSecurityNumber":          {248, 5, packedNumber, required},
		"DateBirth":                     {253, 5, packedDate, required},
		"TelephoneNumber":               {258, 6, packedNumber, nullable},
		"ECOACode":                      {264, 1, alphanumeric, required},
		"ConsumerInformationIndicator":  {265, 2, alphanumeric, applicable},
		"CountryCode":                   {267, 2, alphanumeric, nullable},
		"FirstLineAddress":              {269, 32, alphanumeric, required},
		"SecondLineAddress":             {301, 32, alphanumeric, applicable},
		"City":                          {333, 20, alphanumeric, required},
		"State":                         {353, 2, alphanumeric, required},
		"ZipCode":                       {355, 9, alphanumeric, required},
		"AddressIndicator":              {364, 1, alphanumeric, nullable},
		"ResidenceCode":                 {365, 1, alphanumeric, nullable},
	}
	headerRecordCharacterFormat = map[string]field{
		"BlockDescriptorWord":         {0, 4, numeric | omitted, applicable},
		"RecordDescriptorWord":        {0, 4, numeric, required},
		"RecordIdentifier":            {4, 6, alphanumeric, required},
		"CycleIdentifier":             {10, 2, alphanumeric, applicable},
		"InnovisProgramIdentifier":    {12, 10, alphanumeric, applicable},
		"EquifaxProgramIdentifier":    {22, 10, alphanumeric, applicable},
		"ExperianProgramIdentifier":   {32, 5, alphanumeric, applicable},
		"TransUnionProgramIdentifier": {37, 10, alphanumeric, applicable},
		"ActivityDate":                {47, 8, numeric, required},
		"DateCreated":                 {55, 8, numeric, required},
		"ProgramDate":                 {63, 8, numeric, nullable},
		"ProgramRevisionDate":         {71, 8, numeric, nullable},
		"ReporterName":                {79, 40, alphanumeric, required},
		"ReporterAddress":             {119, 96, alphanumeric, required},
		"ReporterTelephoneNumber":     {215, 10, numeric, nullable},
		"SoftwareVendorName":          {225, 40, alphanumeric, applicable},
		"SoftwareVersionNumber":       {265, 5, alphanumeric, applicable},
		"PRBCProgramIdentifier":       {270, 10, alphanumeric, applicable},
		"Reserved":                    {280, 146, alphanumeric, nullable},
	}
	headerRecordPackedFormat = map[string]field{
		"BlockDescriptorWord":         {0, 4, descriptor | omitted, applicable},
		"RecordDescriptorWord":        {0, 4, descriptor, required},
		"RecordIdentifier":            {4, 6, alphanumeric, required},
		"CycleIdentifier":             {10, 2, alphanumeric, applicable},
		"InnovisProgramIdentifier":    {12, 10, alphanumeric, applicable},
		"EquifaxProgramIdentifier":    {22, 10, alphanumeric, applicable},
		"ExperianProgramIdentifier":   {32, 5, alphanumeric, applicable},
		"TransUnionProgramIdentifier": {37, 10, alphanumeric, applicable},
		"ActivityDate":                {47, 8, numeric, required},
		"DateCreated":                 {55, 8, numeric, required},
		"ProgramDate":                 {63, 8, numeric, nullable},
		"ProgramRevisionDate":         {71, 8, numeric, nullable},
		"ReporterName":                {79, 40, alphanumeric, required},
		"ReporterAddress":             {119, 96, alphanumeric, required},
		"ReporterTelephoneNumber":     {215, 10, numeric, nullable},
		"SoftwareVendorName":          {225, 40, alphanumeric, applicable},
		"SoftwareVersionNumber":       {265, 5, alphanumeric, applicable},
		"PRBCProgramIdentifier":       {270, 10, alphanumeric, applicable},
		"Reserved":                    {280, 86, alphanumeric, nullable},
	}
	trailerRecordCharacterFormat = map[string]field{
		"RecordDescriptorWord":             {0, 4, numeric, required},
		"RecordIdentifier":                 {4, 7, alphanumeric, required},
		"TotalBaseRecords":                 {11, 9, numeric, required},
		"Reserved1":                        {20, 9, alphanumeric, nullable},
		"TotalStatusCodeDF":                {29, 9, numeric, nullable},
		"TotalConsumerSegmentsJ1":          {38, 9, numeric, applicable},
		"TotalConsumerSegmentsJ2":          {47, 9, numeric, applicable},
		"BlockCount":                       {56, 9, numeric, required},
		"TotalStatusCodeDA":                {65, 9, numeric, nullable},
		"TotalStatusCode05":                {74, 9, numeric, nullable},
		"TotalStatusCode11":                {83, 9, numeric, nullable},
		"TotalStatusCode13":                {92, 9, numeric, nullable},
		"TotalStatusCode61":                {101, 9, numeric, nullable},
		"TotalStatusCode62":                {110, 9, numeric, nullable},
		"TotalStatusCode63":                {119, 9, numeric, nullable},
		"TotalStatusCode64":                {128, 9, numeric, nullable},
		"TotalStatusCode65":                {137, 9, numeric, nullable},
		"TotalStatusCode71":                {146, 9, numeric, nullable},
		"TotalStatusCode78":                {155, 9, numeric, nullable},
		"TotalStatusCode80":                {164, 9, numeric, nullable},
		"TotalStatusCode82":                {173, 9, numeric, nullable},
		"TotalStatusCode83":                {182, 9, numeric, nullable},
		"TotalStatusCode84":                {191, 9, numeric, nullable},
		"TotalStatusCode88":                {200, 9, numeric, nullable},
		"TotalStatusCode89":                {209, 9, numeric, nullable},
		"TotalStatusCode93":                {218, 9, numeric, nullable},
		"TotalStatusCode94":                {227, 9, numeric, nullable},
		"TotalStatusCode95":                {236, 9, numeric, nullable},
		"TotalStatusCode96":                {245, 9, numeric, nullable},
		"TotalStatusCode97":                {254, 9, numeric, nullable},
		"TotalECOACodeZ":                   {263, 9, numeric, nullable},
		"TotalEmploymentSegments":          {272, 9, numeric, nullable},
		"TotalOriginalCreditorSegments":    {281, 9, numeric, nullable},
		"TotalPurchasedToSegments":         {290, 9, numeric, nullable},
		"TotalMortgageInformationSegments": {299, 9, numeric, nullable},
		"TotalPaymentInformationSegments":  {308, 9, numeric, nullable},
		"TotalChangeSegments":              {317, 9, numeric, nullable},
		"TotalSocialNumbersAllSegments":    {326, 9, numeric, nullable},
		"TotalSocialNumbersBaseSegments":   {335, 9, numeric, nullable},
		"TotalSocialNumbersJ1Segments":     {344, 9, numeric, nullable},
		"TotalSocialNumbersJ2Segments":     {353, 9, numeric, nullable},
		"TotalDatesBirthAllSegments":       {362, 9, numeric, nullable},
		"TotalDatesBirthBaseSegments":      {371, 9, numeric, nullable},
		"TotalDatesBirthJ1Segments":        {380, 9, numeric, nullable},
		"TotalDatesBirthJ2Segments":        {389, 9, numeric, nullable},
		"TotalTelephoneNumbersAllSegments": {398, 9, numeric, nullable},
		"Reserved2":                        {407, 19, numeric, nullable},
	}
	trailerRecordPackedFormat = map[string]field{
		"RecordDescriptorWord":             {0, 4, descriptor, required},
		"RecordIdentifier":                 {4, 7, alphanumeric, required},
		"TotalBaseRecords":                 {11, 5, packedNumber, required},
		"Reserved1":                        {16, 5, alphanumeric, nullable},
		"TotalStatusCodeDF":                {21, 5, packedNumber, nullable},
		"TotalConsumerSegmentsJ1":          {26, 5, packedNumber, applicable},
		"TotalConsumerSegmentsJ2":          {31, 5, packedNumber, applicable},
		"BlockCount":                       {36, 5, packedNumber, required},
		"TotalStatusCodeDA":                {41, 5, packedNumber, nullable},
		"TotalStatusCode05":                {46, 5, packedNumber, nullable},
		"TotalStatusCode11":                {51, 5, packedNumber, nullable},
		"TotalStatusCode13":                {56, 5, packedNumber, nullable},
		"TotalStatusCode61":                {61, 5, packedNumber, nullable},
		"TotalStatusCode62":                {66, 5, packedNumber, nullable},
		"TotalStatusCode63":                {71, 5, packedNumber, nullable},
		"TotalStatusCode64":                {76, 5, packedNumber, nullable},
		"TotalStatusCode65":                {81, 5, packedNumber, nullable},
		"TotalStatusCode71":                {86, 5, packedNumber, nullable},
		"TotalStatusCode78":                {91, 5, packedNumber, nullable},
		"TotalStatusCode80":                {96, 5, packedNumber, nullable},
		"TotalStatusCode82":                {101, 5, packedNumber, nullable},
		"TotalStatusCode83":                {106, 5, packedNumber, nullable},
		"TotalStatusCode84":                {111, 5, packedNumber, nullable},
		"TotalStatusCode88":                {116, 5, packedNumber, nullable},
		"TotalStatusCode89":                {121, 5, packedNumber, nullable},
		"TotalStatusCode93":                {126, 5, packedNumber, nullable},
		"TotalStatusCode94":                {131, 5, packedNumber, nullable},
		"TotalStatusCode95":                {136, 5, packedNumber, nullable},
		"TotalStatusCode96":                {141, 5, packedNumber, nullable},
		"TotalStatusCode97":                {146, 5, packedNumber, nullable},
		"TotalECOACodeZ":                   {151, 5, packedNumber, nullable},
		"TotalEmploymentSegments":          {156, 5, packedNumber, nullable},
		"TotalOriginalCreditorSegments":    {161, 5, packedNumber, nullable},
		"TotalPurchasedToSegments":         {166, 5, packedNumber, nullable},
		"TotalMortgageInformationSegments": {171, 5, packedNumber, nullable},
		"TotalPaymentInformationSegments":  {176, 5, packedNumber, nullable},
		"TotalChangeSegments":              {181, 5, packedNumber, nullable},
		"TotalSocialNumbersAllSegments":    {186, 5, packedNumber, nullable},
		"TotalSocialNumbersBaseSegments":   {191, 5, packedNumber, nullable},
		"TotalSocialNumbersJ1Segments":     {196, 5, packedNumber, nullable},
		"TotalSocialNumbersJ2Segments":     {201, 5, packedNumber, nullable},
		"TotalDatesBirthAllSegments":       {206, 5, packedNumber, nullable},
		"TotalDatesBirthBaseSegments":      {211, 5, packedNumber, nullable},
		"TotalDatesBirthJ1Segments":        {216, 5, packedNumber, nullable},
		"TotalDatesBirthJ2Segments":        {221, 5, packedNumber, nullable},
		"TotalTelephoneNumbersAllSegments": {226, 5, packedNumber, nullable},
		"Reserved2":                        {231, 135, alphanumeric, nullable},
	}
)
