package rover

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain: rover", func() {
	Describe("New", func() {
		var (
			mapSize = 10
			rover   = New(mapSize)
		)

		It("should return rover interface", func() {
			Expect(rover).To(BeAssignableToTypeOf(&Rover{}))
		})

		It("should return position = {0, 0}", func() {
			Expect(rover.position.x).To(Equal(0))
			Expect(rover.position.y).To(Equal(0))
		})

		It("should return direction = N", func() {
			Expect(rover.direction).To(Equal(N))
		})
	})

	Describe("GetState", func() {
		var (
			rover = New(10)
		)

		It("should return N:0,0", func() {
			state := rover.GetState()
			Expect(state).To(Equal("N:0,0"))
		})
	})

	Describe("Command", func() {
		var rover *Rover
		BeforeEach(func() {
			rover = New(10)
		})

		Describe("order: F", func() {
			It("should return N:0,1", func() {
				state, err := rover.Command("F")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("N:0,1"))
			})
		})

		Describe("order: F with step", func() {
			It("should return N:0,2", func() {
				state, err := rover.Command("F2")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("N:0,2"))
			})
		})

		Describe("order: B with step", func() {
			It("should return N:0,0", func() {
				rover.position = &position{x: 0, y: 2}
				state, err := rover.Command("B2")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("N:0,0"))
			})
		})

		Describe("order: B", func() {
			It("should return N:0,1", func() {
				rover.position = &position{x: 0, y: 2}
				state, err := rover.Command("B")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("N:0,1"))
			})
		})

		Describe("order: L", func() {
			It("should return W:0,0", func() {
				state, err := rover.Command("L")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("W:0,0"))
			})
		})

		Describe("order: R", func() {
			It("should return E:0,0", func() {
				state, err := rover.Command("R")
				Expect(err).To(BeNil())
				Expect(state).To(Equal("E:0,0"))
			})
		})

		Describe("order: unsupported", func() {
			It("should return N:0,0", func() {
				state, err := rover.Command("ABC")
				Expect(err.Error()).To(Equal("unsupport command: ABC"))
				Expect(state).To(Equal("N:0,0"))
			})
		})

	})

	Describe("turnLeft", func() {
		var (
			rover = New(10)
		)

		When("current direction = N", func() {
			BeforeEach(func() {
				rover.direction = N
			})

			It("should change direction = W", func() {
				rover.turnLeft()
				Expect(rover.direction).To(Equal(W))
			})
		})

		When("current direction = E", func() {
			BeforeEach(func() {
				rover.direction = E
			})

			It("should change direction = N", func() {
				rover.turnLeft()
				Expect(rover.direction).To(Equal(N))
			})
		})

		When("current direction = S", func() {
			BeforeEach(func() {
				rover.direction = S
			})

			It("should change direction = E", func() {
				rover.turnLeft()
				Expect(rover.direction).To(Equal(E))
			})
		})

		When("current direction = W", func() {
			BeforeEach(func() {
				rover.direction = W
			})

			It("should change direction = S", func() {
				rover.turnLeft()
				Expect(rover.direction).To(Equal(S))
			})
		})
	})

	Describe("turnRight", func() {
		var (
			rover = New(10)
		)

		When("current direction = N", func() {
			BeforeEach(func() {
				rover.direction = N
			})

			It("should change direction = E", func() {
				rover.turnRight()
				Expect(rover.direction).To(Equal(E))
			})
		})

		When("current direction = E", func() {
			BeforeEach(func() {
				rover.direction = E
			})

			It("should change direction = S", func() {
				rover.turnRight()
				Expect(rover.direction).To(Equal(S))
			})
		})

		When("current direction = S", func() {
			BeforeEach(func() {
				rover.direction = S
			})

			It("should change direction = W", func() {
				rover.turnRight()
				Expect(rover.direction).To(Equal(W))
			})
		})

		When("current direction = W", func() {
			BeforeEach(func() {
				rover.direction = W
			})

			It("should change direction = N", func() {
				rover.turnRight()
				Expect(rover.direction).To(Equal(N))
			})
		})
	})

	Describe("turnHalfLeft", func() {
		var (
			rover = New(10)
		)

		When("current direction = N", func() {
			BeforeEach(func() {
				rover.direction = N
			})

			It("should change direction = NW", func() {
				rover.turnHalfLeft()
				Expect(rover.direction).To(Equal(NW))
			})
		})

		When("current direction = E", func() {
			BeforeEach(func() {
				rover.direction = E
			})

			It("should change direction = NE", func() {
				rover.turnHalfLeft()
				Expect(rover.direction).To(Equal(NE))
			})
		})

		When("current direction = S", func() {
			BeforeEach(func() {
				rover.direction = S
			})

			It("should change direction = SE", func() {
				rover.turnHalfLeft()
				Expect(rover.direction).To(Equal(SE))
			})
		})

		When("current direction = W", func() {
			BeforeEach(func() {
				rover.direction = W
			})

			It("should change direction = SW", func() {
				rover.turnHalfLeft()
				Expect(rover.direction).To(Equal(SW))
			})
		})
	})

	Describe("turnHalfRight", func() {
		var (
			rover = New(10)
		)

		When("current direction = N", func() {
			BeforeEach(func() {
				rover.direction = N
			})

			It("should change direction = NE", func() {
				rover.turnHalfRight()
				Expect(rover.direction).To(Equal(NE))
			})
		})

		When("current direction = E", func() {
			BeforeEach(func() {
				rover.direction = E
			})

			It("should change direction = SE", func() {
				rover.turnHalfRight()
				Expect(rover.direction).To(Equal(SE))
			})
		})

		When("current direction = S", func() {
			BeforeEach(func() {
				rover.direction = S
			})

			It("should change direction = SW", func() {
				rover.turnHalfRight()
				Expect(rover.direction).To(Equal(SW))
			})
		})

		When("current direction = W", func() {
			BeforeEach(func() {
				rover.direction = W
			})

			It("should change direction = NW", func() {
				rover.turnHalfRight()
				Expect(rover.direction).To(Equal(NW))
			})
		})
	})

	Describe("step", func() {
		var (
			rover *Rover
		)

		When("current direction = N", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = N
			})

			Describe("step = 1", func() {
				It("should move position to {x: 0, y: 1} when current = {x: 0, y: 0}", func() {
					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(1))
				})
			})

			Describe("step = 3", func() {
				It("should move position to {x: 0, y: 3} when current = {x: 0, y: 0}", func() {
					rover.step(3)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(3))
				})
			})

			Describe("step = -1", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(-1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = 10", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(10)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = NE", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = NE
			})

			Describe("step = 1", func() {
				It("should move position to {x: 1, y: 1} when current = {x: 0, y: 0}", func() {
					rover.step(1)
					Expect(rover.position.x).To(Equal(1))
					Expect(rover.position.y).To(Equal(1))
				})
			})
		})

		When("current direction = E", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = E
			})

			Describe("step = 1", func() {
				It("should move position to {x: 1, y: 0} when current = {x: 0, y: 0}", func() {
					rover.step(1)
					Expect(rover.position.x).To(Equal(1))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = 3", func() {
				It("should move position to {x: 3, y: 0} when current = {x: 0, y: 0}", func() {
					rover.step(3)
					Expect(rover.position.x).To(Equal(3))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = -1", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(-1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = 10", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(10)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = SE", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = SE
			})

			Describe("step = 1", func() {
				It("should move position to {x: 1, y: 0} when current = {x: 0, y: 1}", func() {
					rover.position = &position{x: 0, y: 1}

					rover.step(1)
					Expect(rover.position.x).To(Equal(1))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = S", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = S
			})

			Describe("step = 1", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})

				It("should move position to {x: 0, y: 3} when current = {x: 0, y: 4}", func() {
					rover.position = &position{x: 0, y: 4}
					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(3))
				})
			})

			Describe("step = 3", func() {
				It("should move position to {x: 0, y: 1} when current = {x: 0, y: 4}", func() {
					rover.position = &position{x: 0, y: 4}
					rover.step(3)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(1))
				})
			})

			Describe("step = -1", func() {
				It("should move position to {x: 0, y: 1} when current = {x: 0, y: 0}", func() {
					rover.step(-1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(1))
				})
			})

			Describe("step = 10", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(10)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = SW", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = SW
			})

			Describe("step = 1", func() {
				It("should move position to {x: 0, y: 0} when current = {x: 1, y: 1}", func() {
					rover.position = &position{x: 1, y: 1}

					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = W", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = W
			})

			Describe("step = 1", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})

				It("should move position to {x: 3, y: 0} when current = {x: 4, y: 0}", func() {
					rover.position = &position{x: 4, y: 0}
					rover.step(1)
					Expect(rover.position.x).To(Equal(3))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = 3", func() {
				It("should move position to {x: 1, y: 0} when current = {x: 4, y: 0}", func() {
					rover.position = &position{x: 4, y: 0}
					rover.step(3)
					Expect(rover.position.x).To(Equal(1))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = -1", func() {
				It("should move position to {x: 1, y: 0} when current = {x: 0, y: 0}", func() {
					rover.step(-1)
					Expect(rover.position.x).To(Equal(1))
					Expect(rover.position.y).To(Equal(0))
				})
			})

			Describe("step = 10", func() {
				It("should not move when current = {x: 0, y: 0}", func() {
					rover.step(10)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(0))
				})
			})
		})

		When("current direction = NW", func() {
			BeforeEach(func() {
				rover = New(5)
				rover.direction = NW
			})

			Describe("step = 1", func() {
				It("should move position to {x: 0, y: 1} when current = {x: 0, y: 1}", func() {
					rover.position = &position{x: 0, y: 1}

					rover.step(1)
					Expect(rover.position.x).To(Equal(0))
					Expect(rover.position.y).To(Equal(1))
				})
			})
		})
	})
})
