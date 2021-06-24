package main

import "fmt"

//status := []string{"Opened", "Closed", "Unlocked", "Locked"}
//action := []string{"closeDoor", "openDoor", "lockDoor", "unlockDoor"}

//table := map[string]string {
//	status[0]+action[0]:status[1],
//	status[1]+action[1]:status[0],
//	status[1]+action[2]:status[3],
//	status[3]+action[3]:status[2],
//	status[2]+action[1]:status[0],
//	status[2]+action[2]:status[3],
//}

type state interface {
	open() error
	close() error
	lock() error
	unlock() error
	show()
}

//=====================================
type door struct {
	openedstate   state
	closedstate   state
	lockedstate   state
	unlockedstate state

	currentstate state
}

func newdoor() *door {
	d := &door{}

	// 接口是指针类型
	d.openedstate = &openedstate{d}
	d.closedstate = &closedstate{d}
	d.lockedstate = &lockedstate{d}
	d.unlockedstate = &unlockedstate{d}

	d.setstate(d.openedstate)

	return d
}

func (d *door) setstate(s state) {
	d.currentstate = s
}

func (d *door) getstate() {
	d.currentstate.show()
}

func (d *door) open() error {
	return d.currentstate.open()
}

func (d *door) close() error {
	return d.currentstate.close()
}

func (d *door) lock() error {
	return d.currentstate.lock()
}

func (d *door) unlock() error {
	return d.currentstate.unlock()
}

//==============================================
type openedstate struct {
	door *door
}

func (s *openedstate) open() error {
	return fmt.Errorf("Error!")
}

func (s *openedstate) close() error {
	s.door.setstate(s.door.closedstate)
	return nil
}

func (s *openedstate) lock() error {
	return fmt.Errorf("Error!")
}

func (s *openedstate) unlock() error {
	return fmt.Errorf("Error!")
}

func (s *openedstate) show() {
	fmt.Println("OPENED")
}

//==============================================
type closedstate struct {
	door *door
}

func (s *closedstate) open() error {
	s.door.setstate(s.door.openedstate)
	return nil
}

func (s *closedstate) close() error {
	return fmt.Errorf("Error!")
}

func (s *closedstate) lock() error {
	s.door.setstate(s.door.lockedstate)
	return nil
}

func (s *closedstate) unlock() error {
	return fmt.Errorf("Error!")
}

func (s *closedstate) show() {
	fmt.Println("CLOSED")
}

//==============================================
type lockedstate struct {
	door *door
}

func (s *lockedstate) open() error {
	return fmt.Errorf("Error!")
}

func (s *lockedstate) close() error {
	return fmt.Errorf("Error!")
}

func (s *lockedstate) lock() error {
	return fmt.Errorf("Error!")
}

func (s *lockedstate) unlock() error {
	s.door.setstate(s.door.unlockedstate)
	return nil
}

func (s *lockedstate) show() {
	fmt.Println("LOCKED")
}

//==============================================
type unlockedstate struct {
	door *door
}

func (s *unlockedstate) open() error {
	s.door.setstate(s.door.openedstate)
	return nil
}

func (s *unlockedstate) close() error {
	return fmt.Errorf("Error!")
}

func (s *unlockedstate) lock() error {
	s.door.setstate(s.door.lockedstate)
	return nil
}

func (s *unlockedstate) unlock() error {
	return fmt.Errorf("Error!")
}

func (s *unlockedstate) show() {
	fmt.Println("UNLOCKED")
}

func main() {
	door := newdoor()
	door.getstate()
	fmt.Println(door.close())
	door.getstate()
	fmt.Println(door.lock())
	door.getstate()
	fmt.Println(door.unlock())
	door.getstate()
	fmt.Println(door.close())
	door.getstate()
}
