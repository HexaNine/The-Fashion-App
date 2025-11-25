package thefashion.authservice.service.abstraction;

import thefashion.authservice.domain.dto.AuthResponseDTO;
import thefashion.authservice.domain.dto.LoginRequestDTO;
import thefashion.authservice.domain.dto.LoginResponseDTO;
import thefashion.authservice.domain.dto.RegisterRequestDTO;

public interface AuthServiceAbstract {
    AuthResponseDTO register(RegisterRequestDTO registerRequestDTO);
    LoginResponseDTO login(LoginRequestDTO loginRequestDTO);
}
