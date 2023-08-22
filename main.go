package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"

	types "github.com/Vadi/pkg/types"
	"gopkg.in/yaml.v3"
)

// StarInfo is a struct to hold star info
type StarInfo struct {
	CategoryName    string
	SubcategoryName string
	ItemName        string
	Accepted        string
	StarCount       int
	Project         string
}

func meanAndStdDev(starInfos []StarInfo) (float64, float64) {
	var sum int
	for _, starInfo := range starInfos {
		sum += starInfo.StarCount
	}
	mean := float64(sum) / float64(len(starInfos))

	var sqSum float64
	for _, starInfo := range starInfos {
		sqSum += math.Pow(float64(starInfo.StarCount)-mean, 2)
	}
	stdDev := math.Sqrt(sqSum / float64(len(starInfos)))

	return mean, stdDev
}

func main() {
	filename := "processed_landscape.yml"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Failed reading file: %s", err)
	}

	landscapeList := types.ProccessedLandscapeList{}

	err = yaml.Unmarshal(data, &landscapeList)

	if err != nil {
		fmt.Printf("unmarshal failed %v", err)
	} else {
		var starList []StarInfo
		for _, ls := range landscapeList.Landscape {
			for _, sub := range ls.Subcategories {
				for _, subitems := range sub.Items {
					if subitems.GithubData.Stars == 0 {
						continue
					}
					starList = append(starList, StarInfo{
						CategoryName:    ls.Name,
						SubcategoryName: sub.Name,
						ItemName:        subitems.Name,
						Accepted:        subitems.Extra.Accepted,
						StarCount:       subitems.GithubData.Stars,
						Project:         subitems.Project,
					})
				}
			}
		}

		sort.Slice(starList, func(i, j int) bool {
			return starList[i].StarCount > starList[j].StarCount
		})

		mean, stdDev := meanAndStdDev(starList)
		thresold := mean + stdDev

		fmt.Println(fmt.Sprintf("Projects rated above thresold: %d", int(thresold)))

		file, err := os.Create("top_hotels_by_thresold.csv")

		if err != nil {
			log.Fatalf("Failed creating file: %s", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"CatgoryName", "SubcategoryName", "ItemName", "Accepted", "StarCount", "Project"})

		for _, kv := range starList {
			writer.Write([]string{
				kv.CategoryName,
				kv.SubcategoryName,
				kv.ItemName,
				kv.Accepted,
				fmt.Sprintf("%d", kv.StarCount),
				kv.Project,
			})
		}

		/*
			percentile := 0.10
			cutOffIndex := int(float64(len(starList)) * percentile)

			fmt.Println(fmt.Sprintf("Top %d percent projects", int(percentile*100)))
			for i := 0; i < cutOffIndex; i++ {
				fmt.Printf("%s: %d\n", starList[i].Name, starList[i].StarCount)
			}
		*/
	}
}
