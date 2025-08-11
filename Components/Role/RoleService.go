package Role

import (
	"github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/dto"
)

type RoleService struct {
	roleRepository *RoleRepository
}

func NewRoleService(permissionRepository *RoleRepository) *RoleService {
	return &RoleService{roleRepository: permissionRepository}
}

func (roleService *RoleService) Create(request dto.CreateRoleReq) (response Entity.Role, err error) {

	role, err := roleService.roleRepository.Create(request)
	if err != nil {
		return Entity.Role{}, err
	}

	return role, nil
}

func (roleService *RoleService) GetAll(request dto.GetAllRoleReq) (response dto.GetAllRolesRes, err error) {
	roles, err := roleService.roleRepository.Get(request)
	if err != nil {
		return dto.GetAllRolesRes{}, err
	}

	return dto.GetAllRolesRes{Roles: roles}, nil
}

func (roleService *RoleService) GetRoleById(Id string) (dto.GetRoleRes, error) {
	roles, err := roleService.roleRepository.GetRoleById(Id)
	if err != nil {
		return dto.GetRoleRes{}, err
	}

	return dto.GetRoleRes{Roles: roles}, nil
}

func (roleService *RoleService) Update(request dto.UpdateRoleReq, roleId string) (dto.GetRoleRes, error) {
	role, err := roleService.roleRepository.Update(request, roleId)
	if err != nil {
		return dto.GetRoleRes{}, err
	}

	return dto.GetRoleRes{Roles: role}, nil
}

func (roleService *RoleService) Delete(roleId string) (dto.GetRoleRes, error) {
	role, err := roleService.roleRepository.Delete(roleId)
	if err != nil {
		return dto.GetRoleRes{}, err
	}

	return dto.GetRoleRes{Roles: role}, nil
}
