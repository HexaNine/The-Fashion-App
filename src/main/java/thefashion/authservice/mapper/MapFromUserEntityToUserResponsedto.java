package thefashion.authservice.mapper;

import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import thefashion.authservice.domain.dto.profile.UserProfileResponseDTO;
import thefashion.authservice.domain.entity.UserEntity;

@Mapper(componentModel = "spring")
public interface MapFromUserEntityToUserResponsedto {
    @Mapping(source = "fistName" , target = "firstName")
    @Mapping(source = "role", target = "role")
    UserProfileResponseDTO toResponseUserEntity(UserEntity userEntity);
}
