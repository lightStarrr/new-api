package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/console_setting"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetStatusReturnsSetupGuideVisibility(t *testing.T) {
	setting := console_setting.GetConsoleSetting()
	original := setting.SetupGuideEnabled
	t.Cleanup(func() {
		setting.SetupGuideEnabled = original
	})

	for _, enabled := range []bool{true, false} {
		setting.SetupGuideEnabled = enabled
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		GetStatus(ctx)

		var response struct {
			Data struct {
				SetupGuideEnabled bool `json:"setup_guide_enabled"`
			} `json:"data"`
		}
		require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
		require.Equal(t, enabled, response.Data.SetupGuideEnabled)
	}
}

func TestGetTopUpInfoReturnsDisabledNoticeVisibility(t *testing.T) {
	setting := operation_setting.GetPaymentSetting()
	original := setting.ShowTopupDisabledNotice
	t.Cleanup(func() {
		setting.ShowTopupDisabledNotice = original
	})

	for _, enabled := range []bool{true, false} {
		setting.ShowTopupDisabledNotice = enabled
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)

		GetTopUpInfo(ctx)

		var response struct {
			Data struct {
				ShowTopupDisabledNotice bool `json:"show_topup_disabled_notice"`
			} `json:"data"`
		}
		require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
		require.Equal(t, enabled, response.Data.ShowTopupDisabledNotice)
	}
}
