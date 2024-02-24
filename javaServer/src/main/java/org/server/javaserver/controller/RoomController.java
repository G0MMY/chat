package org.server.javaserver.controller;

import org.server.javaserver.model.Room;
import org.server.javaserver.model.User;
import org.server.javaserver.repository.RoomRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/rooms")
public class RoomController {

    @Autowired
    private RoomRepository roomRepository;

    @PostMapping
    public Room createRoom(@RequestBody Room room) {
        return roomRepository.save(room);
    }

    @PostMapping("/{id}")
    public String joinRoom(@PathVariable Long id, @RequestBody User user) {
        try {
            Room room = roomRepository.findById(id).get();
            room.addUser(user);
            roomRepository.save(room);

            return "Room joined";
        } catch (Exception e) {
            return e.toString();
        }
    }

    @GetMapping("/{id}")
    public List<String> getRoomUsers(@PathVariable Long id) {
        try {
            Room room = roomRepository.findById(id).get();

            return room.getUsers().stream().map(User::getUsername).collect(Collectors.toList());
        } catch (Exception e) {
            return null;
        }
    }

    @GetMapping("/user/{username}")
    public List<Room> getRoomUsers(@PathVariable String username) {
        try {
            List<Room> rooms = roomRepository.findAll();

            return rooms.stream().filter(room -> room.isUserInRoom(username)).collect(Collectors.toList());
        } catch (Exception e) {
            return null;
        }
    }
}
