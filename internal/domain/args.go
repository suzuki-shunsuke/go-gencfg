package domain

type (
	// InitGenCfgFileArgs is an argument of usecase.InitGenCfgFile .
	InitGenCfgFileArgs struct {
		Dest     string
		TmplDest string
		PNames   []string
		Renderer TemplateRenderer
		Writer   FileWriter
		Exist    FileExist
	}

	// GenCfgFileArgs is an argument of usecase.GenCfgFile .
	GenCfgFileArgs struct {
		Src       string
		Dest      string
		TmplPath  string
		Reader    FileReader
		CfgReader CfgReader
		Generator CodeGenerator
		Executer  CmdExecuter
		CfgUC     CfgUsecase
		EnvUC     EnvUsecase
		FlagUC    FlagUsecase
		ParamUC   ParamUsecase
	}

	// CompareArgs is an argument of usecase.Compare .
	CompareArgs struct {
		Src          string
		Dest         string
		TmplPath     string
		IsQuiet      bool
		IsFailure    bool
		Reader       FileReader
		Renderer     TemplateRenderer
		CfgReader    CfgReader
		Executer     CmdExecuter
		StrFormatter StrFormatter
		CfgUC        CfgUsecase
		EnvUC        EnvUsecase
		FlagUC       FlagUsecase
		ParamUC      ParamUsecase
	}
)
