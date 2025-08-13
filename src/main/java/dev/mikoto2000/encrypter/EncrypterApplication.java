package dev.mikoto2000.encrypter;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

@SpringBootApplication
public class EncrypterApplication {

  public static void main(String[] args) {
    try (ConfigurableApplicationContext ctx = SpringApplication.run(EncrypterApplication.class, args)) {
      BCryptEncrypter encoder = ctx.getBean(BCryptEncrypter.class);
      System.out.println(encoder.encode(args[0]));
    }
  }
}
