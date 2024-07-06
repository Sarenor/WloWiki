import { writable } from 'svelte/store';
	
export let studentsArray = writable([
	{
		"Name" : "Maria",
		"Favorite Subject" : "Math",
		"Age" : 14
	},
	{
		"Name" : "Jose",
		"Favorite Subject" : "Science",
		"Age" : 13
	},
]);

export let colors = writable([
	{
		"Name" : "Cyan",
		"HEX1" : "#00FFFF"
	},
	{
		"Name" : "Yellow",
		"HEX1" : "#FFFF00"
	},
	{
		"Name" : "Blue",
		"HEX1" : "#0000FF"
	},
	{
		"Name" : "Lime",
		"HEX1" : "#00FF00"
	},
	{
		"Name" : "Red",
		"HEX1" : "#FF0000"
	},
	{
		"Name" : "White",
		"HEX1" : "#FFFFFF"
	},
	{
		"Name" : "Black",
		"HEX1" : "#000000"
	},
])

