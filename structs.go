package main

type Design struct {
	Positions [][]string
}

type Player struct {
	Notation int
	Symbol   string
}

const emptySlot = "-"

var playType int
var gridSize int
var freeSlots []int
var player1, player2 Player
var design Design
var currentPlayer Player
