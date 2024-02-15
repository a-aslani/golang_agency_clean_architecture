package discovery_session_request

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/contract/mocks"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/enum"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/notification"
	mocktelegram "github.com/a-aslani/golang_agency_clean_architecture/pkg/notification/mocks"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha"
	mockrecaptcha "github.com/a-aslani/golang_agency_clean_architecture/pkg/recaptcha/mocks"
	"github.com/go-faker/faker/v4"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestRunDiscoverySessionCreateInteractor_Execute(t *testing.T) {

	cfg, err := configs.InitConfig("../../../../config.test.yml")
	if err != nil {
		t.Fatalf("reading config file error: %v", err)
	}

	files := make([]string, 0)

	count := 3

	for i := 0; i < count; i++ {
		files = append(files, uuid.New().String())
	}

	filesObjs := make([]*entity.File, 0)

	for _, idStr := range files {
		id, err := vo.NewFileID(idStr)
		require.NoError(t, err)
		filesObjs = append(filesObjs, &entity.File{
			ID:      id,
			Name:    vo.FileName(faker.Name()),
			Path:    vo.FilePath(faker.URL()),
			Created: time.Now(),
		})
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockrepo := mocks.NewMockRepository(ctrl)

	mockrepo.EXPECT().FindFilesByIDs(gomock.Any(), files).Times(1).Return(filesObjs, nil)

	req := entity.DiscoverySessionCreateRequest{
		ID:             uuid.New().String(),
		Now:            time.Now(),
		Name:           vo.DiscoverySessionName(faker.Name()),
		Email:          vo.DiscoverySessionEmail(faker.Email()),
		ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
		Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
		Files:          filesObjs,
	}

	discoverySessionObj, err := entity.NewDiscoverySession(req)

	mockrepo.EXPECT().SaveDiscoverySession(gomock.Any(), discoverySessionObj).Times(1).Return(nil)

	recaptchaMock := mockrecaptcha.NewMockRecaptcha(ctrl)

	recaptchaToken := uuid.New().String()
	secret := uuid.New().String()

	recaptchaMock.EXPECT().SiteVerify(gomock.Any(), secret, recaptchaToken).Times(1).Return(nil)

	roleID1, err := vo.NewRoleID(uuid.New().String())
	require.NoError(t, err)

	roleID2, err := vo.NewRoleID(uuid.New().String())
	require.NoError(t, err)

	roles := []*entity.Role{
		{
			ID:   roleID1,
			Code: vo.RoleCode(uuid.New().String()),
			Name: vo.RoleName(faker.Name()),
		},
		{
			ID:   roleID2,
			Code: vo.RoleCode(uuid.New().String()),
			Name: vo.RoleName(faker.Name()),
		},
	}

	mockrepo.EXPECT().FindRolesByCodes(gomock.Any(), []string{enum.CEO, enum.CTO}).Times(1).Return(roles, nil)

	chatIds := []int64{faker.RandomUnixTime()}

	mockrepo.EXPECT().FindChatIdsByRoles(gomock.Any(), roles).Times(1).Return(chatIds, nil)

	telegramBotMock := mocktelegram.NewMockTelegramBot(ctrl)

	_, err = generateMessage(discoverySessionObj, filesObjs, cfg.APIUrl)
	require.NoError(t, err)

	for _, chatId := range chatIds {
		telegramBotMock.EXPECT().SendMessage(gomock.Any(), chatId, gomock.Any(), tgbotapi.ModeMarkdown).Times(1).Return(nil)
	}

	usecase := NewUsecase(struct {
		contract.Repository
		recaptcha.Recaptcha
		notification.TelegramBot
	}{
		mockrepo,
		recaptchaMock,
		telegramBotMock,
	})

	res, err := usecase.Execute(context.Background(), InportRequest{
		UUID:           req.ID,
		Now:            req.Now,
		Name:           req.Name,
		Email:          req.Email,
		ProjectDetails: req.ProjectDetails,
		Date:           req.Date,
		RecaptchaToken: recaptchaToken,
		Secret:         secret,
		Files:          files,
	})

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.ID, req.ID)
	require.NotEmpty(t, res.Message)
}
