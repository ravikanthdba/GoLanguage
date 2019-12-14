/*
	 TODO: give "D. Buggs" full gear
 */

/*
	Algorithm:
		1) create a new slice
		2) iterate over the agents slice
		3) find all the agents whose name is not D.buggs
		4) Append all of them to the new slice
		5) find the agent whose name is D.buggs
		6) For the key name: D. buggs, set the gear to full
		7) append the new key-value for D. buggs to the new slice
		8) re-assign new slice data to the original slice
 */

package main

func main() {
	agents := make([]Agent, 0)
	agents = append(agents, Agent{name: "J. Son", gear: "full"})
	agents = append(agents, Agent{name: "A. Pend", gear: "full"})
	agents = append(agents, Agent{name: "D. Buggs", gear: "none"})
	agents = append(agents, Agent{name: "X. Itwon", gear: "full"})
	agents = append(agents, Agent{name: "D. Fercloze", gear: "full"})


	// TODO: give "D. Buggs" full gear
	// NOTE: Output is validated, so don't use extra print statements

	newagents := make([]Agent, 0)
	for _, agent := range agents {
		if agent.name != "D. Buggs" {
			newagents = append(newagents, agent)
		}

		if agent.name == "D. Buggs" {
			newagents = append(newagents, Agent{
				name: "D. Buggs",
				gear: "full",
			})
		}
	}

	agents = newagents

	println("Operation Go: Agent Manifest")
	println("----------------------------")
	for _, agent := range agents {
		println(agent.name, "-> Gear:", agent.gear)
	}
}

// Agent represents an agency employee
type Agent struct {
	name string
	gear string
}