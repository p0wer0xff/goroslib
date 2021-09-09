//nolint:golint,lll
package vision_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type BoundingBox3DArray struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Boxes       []BoundingBox3D
}
