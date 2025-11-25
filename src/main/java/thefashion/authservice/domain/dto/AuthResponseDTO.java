package thefashion.authservice.domain.dto;

import lombok.*;

@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
@Builder
public class AuthResponseDTO {
    private String token;
    private String userId;
    private String firstName;
    private String lastName;
    private String email;
    private String role;
}
