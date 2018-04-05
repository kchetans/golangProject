package config

type JBNSetup struct {
	MongoDBHosts string
	Database     string
	UserName     string
	Password     string
	S3Bucket     string
	S3Secret     string
}

type AWSS3setup struct{
	AWSAccessKeyID 		string
	AWSSecretAccessKey	string
	Token				string
	Bucket				string
}

type SmsSetup struct{ 
	ExotelSid string
	ExotelToken string
	}
func LocalJBNSetup() JBNSetup {
	return JBNSetup{
		MongoDBHosts: "localhost:27017",
		Database:     "jbnlocaldb",
		UserName:     "jbnlocal",
		Password:     "jbnl12345$",
		S3Bucket:     "",
		S3Secret:     "",
	}
}

func DevJBNSetup() JBNSetup {
	return JBNSetup{
		MongoDBHosts: "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318",
		Database:     "jbndevdb",
		UserName:     "jbndev",
		Password:     "jbnd12345$",
		S3Bucket:     "",
		S3Secret:     "",
	}
}

func DemoJBNSetup() JBNSetup {
	return JBNSetup{
		MongoDBHosts: "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318",
		Database:     "jbndemodb",
		UserName:     "jbndemo",
		Password:     "jbnm12345$",
		S3Bucket:     "",
		S3Secret:     "",
	}
}

func QAJBNSetup() JBNSetup {
	return JBNSetup{
		MongoDBHosts: "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318",
		Database:     "jbnqadb",
		UserName:     "jbnqa",
		Password:     "jbnq12345$",
		S3Bucket:     "",
		S3Secret:     "",
	}
}

func ProductionJBNSetup() JBNSetup {
	return JBNSetup{
		MongoDBHosts: "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318",
		Database:     "jbnproddb",
		UserName:     "jbnprod",
		Password:     "jbnp12345$",
		S3Bucket:     "",
		S3Secret:     "",
	}
}


func AmazonS3Bucket() AWSS3setup{
	return AWSS3setup{
		AWSAccessKeyID : "AKIAI2SR75Z7FRNVRG4Q",
		AWSSecretAccessKey : "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",
		Token : "",
		Bucket :"joybynaturedev",
	
	}
}

func OtpSms() SmsSetup{
	return SmsSetup{ 
		ExotelSid : "joybynature",
		ExotelToken : "f590a926362c6cb3043fa45002fbd9c2ed40ee21",
	}
}