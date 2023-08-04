package huddle

type Huddle[T any] struct {

}

type Member[T any] struct {

}

func (h *Huddle[T]) Join() *Member[T] {

}

func (h *Member[T]) Listen() <-chan T {

}

func (h *Member[T]) Say(msg T) {
}

func (h *Member[T]) NothingToSay() {

}


