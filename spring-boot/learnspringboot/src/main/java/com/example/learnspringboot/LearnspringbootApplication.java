package com.example.learnspringboot;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.repository.ListCrudRepository;
import org.springframework.web.servlet.function.RouterFunction;
import org.springframework.web.servlet.function.ServerResponse;

import static org.springframework.web.servlet.function.RouterFunctions.route;
import static org.springframework.web.servlet.function.ServerResponse.ok;

@SpringBootApplication
public class LearnspringbootApplication {

	public static void main(String[] args) {
		java.util.TimeZone.setDefault(java.util.TimeZone.getTimeZone("UTC"));
		SpringApplication.run(LearnspringbootApplication.class, args);
	}

	@Bean
	RouterFunction<ServerResponse> myRoutes(CustomerRepository repository) {
		return route().GET("/customers", _ -> ok().body(repository.findAll())).build();
	}

}

record Customer(int id, String name) {
}

interface CustomerRepository extends ListCrudRepository<Customer, Integer> {

}
