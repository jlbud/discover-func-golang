package main

import (
	"fmt"
	"testing"
)

const (
	readyStartState = iota + 1
	collectingState
	completeState
	suspendedState
	endState
	deleteState
)

func operate(inputState int) {
	switch inputState {
	case readyStartState:
		fmt.Println(readyStartState)
	case collectingState:
		fmt.Println(collectingState)
	case completeState:
		fmt.Println(completeState)
	case suspendedState:
		fmt.Println(suspendedState)
	case endState:
		fmt.Println(endState)
	case deleteState:
		fmt.Println(deleteState)
	}
}

func TestOutputState(t *testing.T) {
	fmt.Println("readyStartState", readyStartState)
	fmt.Println("deleteState", deleteState)
}

func TestOperateState(t *testing.T) {
	operate(1)
}

func TestIota1(t *testing.T) {
	type Type int

	const (
		Bei                     Type = iota + 1 // 素质测评
		Personality                             // 性格测评
		Skill                                   // 专业技能测评
		Knowledge                               // 专业知识测评
		Potential                               // 潜力测评
		WorkValues                              // 工作选择价值观测评
		KeyExp                                  // 关键经历测评
		EmotionalIntelligence                   // 情绪智力测评
		CriticalThinking                        // 批判思维测评
		PracticalIntelligence                   // 管理实践能力测评
		OccupationalPersonality                 // 职业人格测评
		PersonalityDisorder                     // 偏离因素测评
		LeadershipStyle                         // 领导风格测评
		OrgCommitment                           // 组织忠诚度测评
		_                                       // 省略当前项
		_                                       // 省略当前项
		_                                       // 省略当前项
		_                                       // 省略当前项
		_                                       // 省略当前项
		_                                       // 省略当前项
		ManagementQuality                       // 管理素质
	)
	fmt.Println(Bei)
	fmt.Println(KeyExp)
	fmt.Println(OrgCommitment)
	fmt.Println(ManagementQuality)
}
