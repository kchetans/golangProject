package util

type CODETYPE string

//Error Codes
const (
	CODE_IF101   CODETYPE = "IF101"    //infrastructure
	CODE_SR201   CODETYPE = "SR201"    //Server
	CODE_DB301   CODETYPE = "DB301"    //mongo database
	CODE_ST401   CODETYPE = "S-ST401"  //Stories add
	CODE_ST402   CODETYPE = "E-ST402"  //Stories no able to add
	CODE_ST403   CODETYPE = "S-ST403"  //Related Stories found
	CODE_ST404   CODETYPE = "E-ST404"  //Related Stories not found
	CODE_ST405   CODETYPE = "S-ST405"  //
	CODE_ST406   CODETYPE = "E-ST406"  //
	CODE_ST407   CODETYPE = "S-ST407"  //
	CODE_ST408   CODETYPE = "E-ST408"  //
	CODE_ST409   CODETYPE = "S-ST409"  //
	CODE_ST410   CODETYPE = "E-ST410"  //
	CODE_ST411   CODETYPE = "S-ST411"  //
	CODE_ST412   CODETYPE = "E-ST412"  //
	CODE_ST413   CODETYPE = "E-ST413"  //
	CODE_ST414   CODETYPE = "E-ST414"  //
	CODE_ST415   CODETYPE = "E-ST415"  // Get Emox for a story successfully
	CODE_ST416   CODETYPE = "E-ST416"  // Error in finding emoxs for story
	CODE_SP101   CODETYPE = "S-SP101"  //Add spam report successfully
	CODE_SP102   CODETYPE = "E-SP102"  //Error adding spam report
	CODE_SP201   CODETYPE = "S-SP201"  //Get spam report list successfully
	CODE_SP202   CODETYPE = "E-SP202"  //Error getting spam report list
	CODE_QA501   CODETYPE = "S-QA501"  //Question Anwsers
	CODE_QA502   CODETYPE = "E-QA502"  //Question Anwsers
	CODE_QA503   CODETYPE = "E-QA503"  //Add QA Emox
	CODE_QA504   CODETYPE = "E-QA504"  //List QA Emox
	CODE_QA101   CODETYPE = "S-QA101"  //Add qa comments successfully
	CODE_QA102   CODETYPE = "E-QA102"  //Error adding QA comments
	CODE_QA103   CODETYPE = "S-QA103"  //Get QA details successfully
	CODE_QA104   CODETYPE = "E-QA104"  //Error getting QA detail
	CODE_QA105   CODETYPE = "S-QA105"  //Get related question successfully
	CODE_QA106   CODETYPE = "E-QA106"  //Error getting related questions
	CODE_SH601   CODETYPE = "SH601"    //Shop
	CODE_TP701   CODETYPE = "TP701"    //Tip
	CODE_RC801   CODETYPE = "RC801"    //Recipe
	CODE_DT901   CODETYPE = "DT901"    //Diet
	CODE_USR1001 CODETYPE = "USR1001"  //User
	CODE_OTP1101 CODETYPE = "OTP1101"  //Otp InActive
	CODE_OTP2101 CODETYPE = "OTP2101"  //Otp Active
	CODE_TAG001  CODETYPE = "S-TAGP01" //tag found - success
	CODE_TAG002  CODETYPE = "E-TAGP02" // Tag not found
	CODE_USR101  CODETYPE = "S-USR101" // Add User successfully
	CODE_USR102  CODETYPE = "E-USR102" // Error adding user
	CODE_USR103  CODETYPE = "S-USR103" // Login user successfully
	CODE_USR104  CODETYPE = "E-USR104" // Error login user
	CODE_USR105  CODETYPE = "S-USR105" //
	CODE_USR106  CODETYPE = "E-USR106" //User Already Exists
	CODE_USR107  CODETYPE = "E-USR107" //Error reading signup form
	CODE_USR108  CODETYPE = "E-USR108" //Signup form Validation Error
)
