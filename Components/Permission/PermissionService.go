package Permission

import (
	"github.com/mahdidl/golang_boilerplate/dto"
)

type PermissionService struct {
	permissionRepository *PermissionRepository
}

func NewPermissionService(permissionRepository *PermissionRepository) *PermissionService {
	return &PermissionService{permissionRepository: permissionRepository}
}

func (permissionService PermissionService) CreateNewPermission(request dto.CreatePermissionReq) (response dto.CreatePermissionRes, err error) {

	permission, err := permissionService.permissionRepository.CreateNewPermission(request)
	if err != nil {
		return dto.CreatePermissionRes{}, err
	}

	return dto.CreatePermissionRes{ID: permission.Id, Name: permission.Name}, nil
}

func (permissionService PermissionService) GetPermissions(request dto.GetAllPermissionsReq) (response dto.GetPermissionsRes, err error) {
	permissions, err := permissionService.permissionRepository.GetPermissions(request)
	if err != nil {
		return dto.GetPermissionsRes{}, err
	}

	return dto.GetPermissionsRes{Permissions: permissions}, nil
}
