package domain

type (
	// InitGenCfgFileArgs is an argument of usecase.InitGenCfgFile .
	InitGenCfgFileArgs struct {
		Dest     string
		PNames   []string
		Renderer TemplateRenderer
		Writer   FileWriter
		Exist    FileExist
	}

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
		ParamUC      ParamUsecase
	}
)
