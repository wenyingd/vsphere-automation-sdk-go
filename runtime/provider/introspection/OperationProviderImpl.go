package introspection
//
//import (
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
//	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/provider/introspection/bindings/operation"
//)
//
//type OperationImpl struct {
//	introspector *LocalProviderIntrospector
//}
//
//
//func NewOperationServiceImpl(introspector *LocalProviderIntrospector) *OperationImpl{
//    return &OperationImpl{introspector:introspector}
//}
//
//func(OperationImpl OperationImpl) List(serviceId string, ctx *core.ExecutionContext) (map[string]bool, error) {
//	return OperationImpl.introspector.GetOperations(serviceId)
//
//}
//func(OperationImpl OperationImpl) Get(serviceId string, operationId string, ctx *core.ExecutionContext) (operation.Info, error) {
//	return OperationImpl.introspector.GetOperationInfo(serviceId, operationId)
//
//}
//