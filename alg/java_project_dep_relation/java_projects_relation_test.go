package java_project_dep_relation_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

const projects = `com_example_abs:abs-api:jar:2_0_0-SNAPSHOT;com_example_acs:acs-api:jar:1_0_0-SNAPSHOT
com_example_abs:abs-api:jar:2_0_0-SNAPSHOT;com_example_app-base:app-base-common:jar:2_0_2-SNAPSHOT
com_example_abs:abs-api:jar:2_0_0-SNAPSHOT;com_example_app-base:app-base-util:jar:2_0_2-SNAPSHOT
com_example_abs:abs-api:jar:2_0_0-SNAPSHOT;com_example_bmbs:bmbs-service-api:jar:0_1-SNAPSHOT
com_example_abs:abs-api:jar:2_0_0-SNAPSHOT;com_example_ubs:hk-ubs-api:jar:2_0_0-SNAPSHOT`

type Project struct {
	name           string
	level          int
	depProjects    map[string]struct{}
	followProjects map[string]struct{}
}

func (p *Project) addDep(name string) {
	p.depProjects[name] = struct{}{}
}
func (p *Project) addFollow(name string) {
	p.followProjects[name] = struct{}{}
}

func TestProjectOrder(t *testing.T) {
	m := buildMap()
	loopM := checkLoopDep(buildMap())
	if len(loopM) > 0 {
		fmt.Println("----------------------------------")
		fmt.Println("存在循环依赖的项目dot图")
		fmt.Println("----------------------------------")
		fmt.Printf("digraph G {\n")
		for _, p := range m {
			for f := range p.followProjects {
				fmt.Printf("\"%s\" -> \"%s\";\n", f, p.name)
			}
		}
		fmt.Printf("}\n\n")

		fmt.Println("----------------------------------")
		fmt.Println("存在循环依赖的项目")
		fmt.Println("----------------------------------")
		for name := range loopM {
			fmt.Println(name)
			delete(m, name)
			for _, p := range m {
				delete(p.followProjects, name)
				delete(p.depProjects, name)
			}
		}
	}

	calcLevels(m)

}

func buildMap() map[string]*Project {

	m := make(map[string]*Project)

	orders := strings.Split(projects, "\n")
	for _, order := range orders {
		arr := strings.Split(order, ";")
		name1 := getProjectName(arr[0])
		name2 := getProjectName(arr[1])

		p1, ok := m[name1]
		if !ok {
			p1 = &Project{
				name:           name1,
				depProjects:    make(map[string]struct{}),
				followProjects: make(map[string]struct{}),
			}
			m[name1] = p1
		}
		p2, ok := m[name2]
		if !ok {
			p2 = &Project{
				name:           name2,
				depProjects:    make(map[string]struct{}),
				followProjects: make(map[string]struct{}),
			}
			m[name2] = p2
		}

		p1.addDep(name2)
		p2.addFollow(name1)
	}
	return m
}

func getProjectName(s string) string {
	s = strings.ReplaceAll(s, ":compile", "")
	idx := strings.LastIndex(s, ":")
	s = s[:idx]
	idx = strings.LastIndex(s, ":")
	s = s[:idx]
	return s
}

func checkLoopDep(m map[string]*Project) map[string]*Project {
	// 剔除出度为0的项目
	for len(m) > 0 {
		var zeroOutProjects []*Project
		for _, p := range m {
			if len(p.followProjects) == 0 {
				zeroOutProjects = append(zeroOutProjects, p)
			}
		}

		if len(zeroOutProjects) > 0 {
			for _, p := range zeroOutProjects {
				for depProj := range p.depProjects {
					delete(m[depProj].followProjects, p.name)
				}
				delete(m, p.name)

				for _, pp := range m {
					delete(pp.depProjects, p.name)
					delete(pp.followProjects, p.name)
				}
			}
			continue
		}
		break
	}

	// 剔除入度为0的项目
	for len(m) > 0 {
		var zeroOutProjects []*Project
		for _, p := range m {
			if len(p.depProjects) == 0 {
				zeroOutProjects = append(zeroOutProjects, p)
			}
		}

		if len(zeroOutProjects) > 0 {
			for _, p := range zeroOutProjects {
				for followProj := range p.followProjects {
					delete(m[followProj].depProjects, p.name)
				}
				delete(m, p.name)

				for _, pp := range m {
					delete(pp.depProjects, p.name)
					delete(pp.followProjects, p.name)
				}
			}
			continue
		}
		break
	}

	return m

}

func calcLevels(m map[string]*Project) {
	levelMap := make(map[string]*Project)
	for _, p := range m {
		if len(p.depProjects) == 0 {
			levelMap[p.name] = p
		}
	}

	level := 0

	for len(levelMap) > 0 {
		nextLevelMap := make(map[string]*Project)
		for _, p := range levelMap {
			p.level = level
			for fname := range p.followProjects {
				nextLevelMap[fname] = m[fname]
			}
			//fmt.Printf("%5d  %100s  %s\n", p.level, p.name, p.followProjects)
		}
		level++
		levelMap = nextLevelMap
	}

	var projectArr []*Project
	for _, p := range m {
		projectArr = append(projectArr, p)
	}

	sort.Slice(projectArr, func(i, j int) bool {
		p1 := projectArr[i]
		p2 := projectArr[j]
		if p1.level == p2.level {
			return strings.Compare(p1.name, p2.name) < 0
		}
		return p1.level < p2.level
	})

	fmt.Printf("项目名称              部署顺序\n")
	for _, p := range projectArr {
		fmt.Printf("%-60s %d\n", p.name, p.level)
	}
}
