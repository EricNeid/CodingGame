package main

import "fmt"
import "os"
import "bufio"

//import "strings"
//import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)

	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	for i := 0; i < height; i++ {
		scanner.Scan()
		//line := scanner.Text()
	}
	// sanityLossLonely: how much sanity you lose every turn when alone, always 3 until wood 1
	// sanityLossGroup: how much sanity you lose every turn when near another player, always 1 until wood 1
	// wandererSpawnTime: how many turns the wanderer take to spawn, always 3 until wood 1
	// wandererLifeTime: how many turns the wanderer is on map after spawning, always 40 until wood 1
	var sanityLossLonely, sanityLossGroup, wandererSpawnTime, wandererLifeTime int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &sanityLossLonely, &sanityLossGroup, &wandererSpawnTime, &wandererLifeTime)

	for {
		// entityCount: the first given entity corresponds to your explorer
		var entityCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &entityCount)

		var myself entity
		var theOthers, wanderers []entity

		for i := 0; i < entityCount; i++ {
			newEntity := entity{}
			scanner.Scan()
			fmt.Sscan(
				scanner.Text(),
				&newEntity.entityType,
				&newEntity.id,
				&newEntity.x,
				&newEntity.y,
				&newEntity.param0,
				&newEntity.param1,
				&newEntity.param2)
			if i == 0 {
				myself = newEntity
			} else if newEntity.entityType == explorer {
				theOthers = append(theOthers, newEntity)
			} else if newEntity.entityType == wanderer {
				wanderers = append(theOthers, newEntity)
			}
		}

		debug(fmt.Sprintf("Currently %d wanderers around", len(wanderers)))

		var nearest entity
		shortestDistance := 99999.0
		for _, e := range theOthers {
			d := getDistance(myself, e)
			if d < shortestDistance {
				shortestDistance = d
				nearest = e
			}
		}

		fmt.Println(fmt.Sprintf("MOVE %d %d", nearest.x, nearest.y))
	}
}
