package users

// // CheckNamespace .
// func CheckNamespace(ctx *gin.Context) {
// 	exi, err := db.ExistsNamespace(ctx.Param("name"))
// 	if err != nil {
// 		apis.InternalServerErr(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]bool{"message": exi})
// }

// func CreateOrg(ctx *gin.Context) {
// 	org := &Org{}
// 	if err := decodeBody(ctx, org); err != nil {
// 		apis.BadRequestErr(ctx, err)
// 		return
// 	}
// 	if err := validate.IsDNS1035Label(org.Name); err != nil {
// 		apis.BadRequestErr(ctx, err)
// 		return
// 	}

// 	if err := db.CreateOrg(org); err != nil {
// 		apis.InternalServerErr(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, org)
// }

// func ListOrgs(ctx *gin.Context) {
// 	os, err := db.ListOrgs(ListOrgOption{})
// 	if err != nil {
// 		apis.InternalServerErr(ctx, err)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, os)
// }

// func GetOrgProfile(ctx *gin.Context) {

// }
