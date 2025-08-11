package Permission

import (
	"log"

	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	Entity "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	"github.com/mahdidl/golang_boilerplate/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PermissionRepository struct {
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (permissionRepository *PermissionRepository) CreateNewPermission(request dto.CreatePermissionReq) (Entity.Permission, error) {
	permission := Entity.Permission{}

	result, err := Config.PermissionCollection.InsertOne(Config.DBContext, Entity.Permission{Id: primitive.NewObjectID(), Name: request.Name})
	if err != nil {
		return Entity.Permission{}, err
	}

	if err = Config.PermissionCollection.FindOne(Config.DBContext, bson.M{"_id": result.InsertedID}).Decode(&permission); err != nil {
		return Entity.Permission{}, err
	}

	return permission, err
}

func (permissionRepository *PermissionRepository) GetPermissions(request dto.GetAllPermissionsReq) ([]Entity.Permission, error) {
	var permissions = make([]Entity.Permission, 0)

	permissionCursor, queryError := Config.PermissionCollection.Find(Config.DBContext, bson.M{}, Helper.NewMongoPaginate(request.Limit, request.Page).GetPaginatedOpts())
	if queryError != nil {
		return nil, queryError
	}

	// decode permission and append to list
	for permissionCursor.Next(Config.DBContext) {
		var permission Entity.Permission
		if err := permissionCursor.Decode(&permission); err != nil {
			log.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
