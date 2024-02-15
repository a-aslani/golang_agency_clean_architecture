package postgres

import (
	"context"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/project/model/vo"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/enum"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func TestGateway_FindFilesByIDs(t *testing.T) {

	files := make([]*entity.File, 0)
	ids := make([]string, 0)

	for i := 0; i < 3; i++ {
		file := saveFile(t)
		files = append(files, file)
		ids = append(ids, file.ID.String())
	}

	objs, err := datasource.FindFilesByIDs(context.Background(), ids)
	require.NoError(t, err)
	require.NotNil(t, objs)
	require.Len(t, objs, len(files))
}

func TestGateway_SaveDiscoverySession(t *testing.T) {

	id, err := vo.NewDiscoverySessionID(uuid.New().String())
	require.NoError(t, err)

	files := make([]*entity.File, 0)

	for i := 0; i < 3; i++ {
		file := saveFile(t)
		files = append(files, file)
	}

	obj := &entity.DiscoverySession{
		ID:             id,
		Name:           vo.DiscoverySessionName(faker.Name()),
		Email:          vo.DiscoverySessionEmail(faker.Email()),
		ProjectDetails: vo.DiscoverySessionProjectDetails(faker.Sentence()),
		Date:           vo.DiscoverySessionDate(time.Now().Add(72 * time.Hour)),
		Files:          files,
		Created:        time.Now(),
		Updated:        time.Now(),
	}

	err = datasource.SaveDiscoverySession(context.Background(), obj)
	require.NoError(t, err)
}

func TestGateway_SaveFilePath(t *testing.T) {
	saveFile(t)
}

func saveFile(t *testing.T) *entity.File {
	id := uuid.New().String()

	obj := entity.File{
		ID:      vo.FileID(id),
		Name:    vo.FileName(faker.Name()),
		Path:    vo.FilePath(faker.URL()),
		Created: time.Now(),
	}

	err := datasource.SaveFilePath(context.Background(), &obj)
	require.NoError(t, err)

	return &obj
}

func saveRole(t *testing.T) *entity.Role {
	id, err := vo.NewRoleID(uuid.New().String())
	require.NoError(t, err)

	obj := &entity.Role{
		ID:   id,
		Code: vo.RoleCode(uuid.New().String()),
		Name: vo.RoleName(faker.Name()),
	}

	err = datasource.SaveRole(context.Background(), obj)
	require.NoError(t, err)
	return obj
}

func TestGateway_SaveRole(t *testing.T) {
	saveRole(t)
}

func TestGateway_FindRolesByCodes(t *testing.T) {
	codes := []string{enum.CEO, enum.SUPPORT}

	for _, code := range codes {
		id, err := vo.NewRoleID(uuid.New().String())
		require.NoError(t, err)

		obj := &entity.Role{
			ID:   id,
			Code: vo.RoleCode(code),
			Name: vo.RoleName(faker.Name()),
		}

		err = datasource.SaveRole(context.Background(), obj)
		require.NoError(t, err)
	}

	roles, err := datasource.FindRolesByCodes(context.Background(), codes)
	require.NoError(t, err)
	require.NotEmpty(t, roles)
	require.Len(t, roles, len(codes))
}

func TestGateway_SaveChatID(t *testing.T) {
	role := saveRole(t)
	saveTelegramChatID(t, role.ID.String())
}

func saveTelegramChatID(t *testing.T, roleId string) *entity.TelegramChatID {

	rndStr := uuid.New().String()

	id, err := vo.NewTelegramChatIDID(rndStr)
	require.NoError(t, err)
	require.Equal(t, rndStr, id.String())

	obj := &entity.TelegramChatID{
		ID:     id,
		ChatID: strconv.Itoa(int(faker.RandomUnixTime())),
		RoleID: roleId,
	}

	err = datasource.SaveTelegramChatID(context.Background(), obj)
	require.NoError(t, err)

	return obj
}

func TestGateway_FindTelegramChatIdsByRoles(t *testing.T) {

	roles := make([]*entity.Role, 0)

	for i := 0; i < 5; i++ {
		roles = append(roles, saveRole(t))
	}

	chatIds := make([]*entity.TelegramChatID, 0)

	for _, role := range roles {
		chatIds = append(chatIds, saveTelegramChatID(t, role.ID.String()))
	}

	chatIdsRes, err := datasource.FindChatIdsByRoles(context.Background(), roles)
	require.NoError(t, err)
	require.NotEmpty(t, chatIdsRes)
	require.Equal(t, len(chatIds), len(chatIdsRes))
}
