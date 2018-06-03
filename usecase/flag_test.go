package usecase_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-gencfg/domain"
	"github.com/suzuki-shunsuke/go-gencfg/usecase"
)

func TestFlagIsBind(t *testing.T) {
	flag := domain.Flag{}
	flagUC := usecase.FlagUsecase{}
	if flagUC.IsBind(flag, nil) {
		t.Fatal("flagUC.IsBind(flag, nil) = true, wanted false")
	}
	if !flagUC.IsBind(flag, ptr.PBool(true)) {
		t.Fatal("flagUC.IsBind(flag, nil) = false, wanted true")
	}
	flag.Bind = ptr.PBool(true)
	if !flagUC.IsBind(flag, nil) {
		t.Fatal("flagUC.IsBind(flag, nil) = false, wanted true")
	}
	flag.Bind = nil
	flag.Name = ptr.PStr("hello")
	if !flagUC.IsBind(flag, nil) {
		t.Fatal("flagUC.IsBind(flag, nil) = false, wanted true")
	}
}
