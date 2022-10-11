package auth

import (
	"gofiber-todo/src/entity"
	"gofiber-todo/src/modules/auth/dto"
	"gofiber-todo/src/utils/bcrypt"
	"gofiber-todo/src/utils/jwt"
	"gofiber-todo/src/utils/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *AuthService
}

func (ctrl *AuthController) getAccessToken(user *entity.User) string {
	return jwt.Generate(&jwt.TokenPayload{
		ID:    user.ID,
		Email: user.Email,
	})
}

// Login godoc
// @Summary      Login
// @Description  login with email and pw
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        LoginDTO   body    dto.LoginDTO  true  "Email and Password"
// @Success      200  	 {object}  	dto.LoginResult
// @Failure      400  	 {object}		response.ErrorResponse
// @Failure      500  	 {object}		response.ErrorResponse
// @Router       /api/auth/login [post]
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	loginDto := new(dto.LoginDTO)
	if err := validator.ParseBodyAndValidate(c, loginDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//find user with email, password
	userEntity := &entity.User{}
	if err := ctrl.service.findUserByEmail(userEntity, loginDto.Email).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	isValidPw, err := bcrypt.Compare(userEntity.Password, loginDto.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	} else if !isValidPw {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "wrong password",
		})
	}

	accessToken := ctrl.getAccessToken(userEntity)

	return c.JSON(fiber.Map{
		"data": dto.LoginResult{
			AccessToken: accessToken,
		},
	})
}

// SignUp godoc
// @Summary      SignUp
// @Description  signup with email and pw
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        SignupDTO   body   dto.SignupDTO  true  "Email and Password"
// @Success      200  	 {object}  	dto.LoginResult
// @Failure      400  	 {object}		response.ErrorResponse
// @Failure      500  	 {object}		response.ErrorResponse
// @Router       /api/auth/signup [post]
func (ctrl *AuthController) SignUp(c *fiber.Ctx) error {
	signUpDto := new(dto.SignupDTO)
	if err := validator.ParseBodyAndValidate(c, signUpDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPw, err := bcrypt.Generate(signUpDto.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	signUpDto.Password = hashedPw

	userEntity := signUpDto.ToEntity()
	if err := ctrl.service.createUser(userEntity).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	accessToken := ctrl.getAccessToken(userEntity)

	return c.JSON(fiber.Map{
		"data": dto.LoginResult{
			AccessToken: accessToken,
		},
	})
}
