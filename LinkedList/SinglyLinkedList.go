package main

import "fmt"
//some edition for doubly linked list
//new edition for circular linked list
type song struct{
	name string
	artist string
	next *song

	previous *song
}
type playlist struct{
	name string
	head *song

	tail *song
	nowPlaying *song
}
func createPlaylist(name string) *playlist{
	return &playlist{
		name:name,
	}
}
func  (p *playlist)addSong(name, artist string) error  {
	s:=&song{
		name:name,
		artist:artist,
	}
	if p.head==nil{
		p.head=s
	} else {
		//for singly linked list

		//currentNode:=p.head
		//for currentNode.next!=nil{
		//	currentNode=currentNode.next
		//}
		//currentNode.next=s

		//for doubly linked list
		currentNode:=p.tail
		currentNode.next=s
		s.previous=p.tail
	}
	s.next=p.head
	p.tail=s
	p.head.previous=p.tail.next
	return nil
}
func (p *playlist) showAllSongs() error {
	currentNode:=p.head
	if currentNode==nil{
		fmt.Println("Playlist is empty")
		return nil
	}

	for currentNode.next!=p.head.previous{
		fmt.Println(currentNode.name,currentNode.artist)
		currentNode=currentNode.next
	}
	return nil
}
func (p *playlist)startPlaying() *song{
	p.nowPlaying=p.head
	return p.nowPlaying
}
func (p *playlist) nextSong() *song{
	p.nowPlaying=p.nowPlaying.next
	return p.nowPlaying
}

func (p *playlist)previousSong() *song{
	p.nowPlaying=p.nowPlaying.previous
	return p.nowPlaying
}

func main(){

	obj := Constructor();
	fmt.Println(obj)




	//playlistName:="MyPlayList"
	//myPlaylist:=createPlaylist(playlistName)
	//fmt.Println("Playlist is created",myPlaylist,"\n")
	//
	//fmt.Println("Adding new songs in playlist...\n\n")
	//myPlaylist.addSong("True Love","Pink")
	//myPlaylist.addSong("Lovestory","taylor Swift")
	//myPlaylist.addSong("story of My life","One Direction")
	//
	//fmt.Println("Showing all songs in playlist")
	//myPlaylist.showAllSongs()
	//fmt.Println()
	//
	//myPlaylist.startPlaying()
	//fmt.Println(myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist,"\n")
	//
	//myPlaylist.nextSong()
	//fmt.Println("Moving to next song")
	//fmt.Println(myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist,"\n")
	//myPlaylist.nextSong()
	//fmt.Println("Moving to next song")
	//fmt.Println(myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist,"\n")
	//myPlaylist.nextSong()
	//fmt.Println("Moving to next song")
	//fmt.Println(myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist,"\n")
	//
	//myPlaylist.previousSong()
	//fmt.Println("Moving to previous song")
	//fmt.Println(myPlaylist.nowPlaying.name,myPlaylist.nowPlaying.artist,"\n")
}