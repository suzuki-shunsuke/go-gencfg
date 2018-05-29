package domain

type (
	// GenCfgFileArgs is an argument of usecase.GenCfgFileArgs .
	GenCfgFileArgs struct {
		Src          string
		Dest         string
		TmplPath     string
		TestTmplPath string
		Reader       FileReader
		CfgReader    CfgReader
		Generater    CodeGenerater
		Executer     CmdExecuter
		CfgUC        CfgUsecase
		EnvUC        EnvUsecase
		FlagUC       FlagUsecase
	}
)
