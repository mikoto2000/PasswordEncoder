package dev.mikoto2000.encrypter;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;

/**
 * BCryptEncrypter
 */
@Component
public class BCryptEncrypter implements Encrypter {
  private static final PasswordEncoder encoder = new BCryptPasswordEncoder();
  @Override
  public String encode(String str) {
    return encoder.encode(str);
  }
}
