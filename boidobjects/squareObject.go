package boidobjects

import ()

/*
//Make sure that any game objects actually implement the object interface
var _ objects.Object = (*Square)(nil)

type Square struct {
	objects.BaseObject
	env              objects.ObjEnviroment
	offset           int
	resize           chan linmath.PSPoint
	windowDim        linmath.PSPoint
	ObjectNum        int
	generator        *rand.Rand
	velocity         [2]float64
	desiredDirection [2]float64
}

func (s *Square) Update(updates int) {
	if s.windowDim == nil {
		//s.env.GetEventCallback()(event.UpdateEvent{Sender: s.ObjectNum, Receiver: -1, EventCode: event.SubscribeEvent_WindowResize, Event: event.SubscribeEvent_WindowResizeEvent{s.resize}})
		s.windowDim = s.env.GetWindowSize()
	}

	//select {
	//case x := <-s.resize:
	//	s.windowDim = x
	//default:
	//}
	x, y, z := s.GetPos()
	const maxSpeed = 200
	const steerStrength = 200
	const wander = .3
	if updates%10 == 0 {
		s.desiredDirection = [2]float64{s.desiredDirection[0] + float64(frand.Intn(100)-50)/75*wander + (.6 * math.Max(math.Abs(math.Min(0, float64(x))), 1)) - (.6 * math.Max(math.Abs(math.Min(0, 3000-float64(x))), 1)), s.desiredDirection[1] + float64(frand.Intn(100)-50)/75*wander + (.6 * math.Max(math.Abs(math.Min(0, float64(y))), 1)) - (.6 * math.Max(math.Abs(math.Min(0, 3000-float64(y))), 1))}
		norm := math.Sqrt(s.desiredDirection[0]*s.desiredDirection[0] + s.desiredDirection[1]*s.desiredDirection[1])
		s.desiredDirection = [2]float64{s.desiredDirection[0] / norm, s.desiredDirection[1] / norm}
	}
	desiredVelocity := [2]float64{s.desiredDirection[0] * maxSpeed, s.desiredDirection[1] * maxSpeed}
	desiredSteeringForce := [2]float64{(desiredVelocity[0] - s.velocity[0]) * steerStrength, (desiredVelocity[1] - s.velocity[1]) * steerStrength}
	acceleration := [2]float64{math.Min(desiredSteeringForce[0], 200), math.Min(desiredSteeringForce[1], 200)}
	s.velocity = [2]float64{math.Min(s.velocity[0]+(acceleration[0]/60), maxSpeed), math.Min(s.velocity[1]+(acceleration[1]/60), maxSpeed)}
	if updates > 100 {
		s.Move(x+float32(s.velocity[0]/60), y+float32(s.velocity[1]/60), z)
	}

}
func (s *Square) Render(input []float32) {
	s.BaseObject.Render(input)
}

func (s *Square) Init(offset int) {
	s.BaseObject.Init()
	s.generator = rand.New(rand.NewSource(int64(offset)))
	s.desiredDirection = [2]float64{1, 1}
	s.resize = make(chan linmath.PSPoint, 1)
	s.windowDim = nil
	s.Move((float32(1080)/650)*float32(offset%650), ((float32(1080) / 650) * float32(int(offset/650))), 50)
	s.Resize(5, 5)
	blue := 0
	red := 0
	green := 0
	if offset > 0 && offset < 200000/3 {
		red = 150
		green = 206
		blue = 180
	} else if offset < (200000/3)*2 {
		red = 255
		green = 236
		blue = 173
	} else {
		red = 255
		green = 111
		blue = 105
	}
	s.Recolor(uint8(red), uint8(green), uint8(blue), 255)
	s.SetTexSize(0, 0)
	s.SetMap(0)
	s.offset = offset
}

func (s *Square) SendEvent(e event.UpdateEvent) {
	if e.EventCode == event.UpdateEvent_NewObject {
		s.ObjectNum = e.Receiver
	}
}

func (s *Square) SetEnviroment(env objects.ObjEnviroment) {
	s.env = env
	//Set Callback
}
*/
