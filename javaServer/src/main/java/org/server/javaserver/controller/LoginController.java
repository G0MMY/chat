package org.server.javaserver.controller;

import org.server.javaserver.model.User;
import org.server.javaserver.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Objects;

@RestController
@RequestMapping("/login")
public class LoginController {

    @Autowired
    private UserRepository userRepository;

    @PostMapping
    public String login(@RequestBody User user) {
        User result = userRepository.findByUsername(user.getUsername());

        if(Objects.isNull(result)) {
            return "invalid username";
        }

        if(result.getPassword().equals(user.getPassword())) {
            return "logged in";
        }

        return "invalid password";
    }
}
