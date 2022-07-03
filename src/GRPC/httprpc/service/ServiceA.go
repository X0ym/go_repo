package service

// GrpcRequest 及其属性均为可导出的
type GrpcRequest struct {
	X, Y int
}

// GrpcReplay 及其属性均为可导出的
type GrpcReplay struct {
	Code int
	Data int
	Msg  string
}

type ServiceA struct{}

func (s *ServiceA) Add(req *GrpcRequest, resp *GrpcReplay) error {
	resp.Data = req.X + req.Y
	resp.Code = 1
	resp.Msg = "成功"
	return nil
}
