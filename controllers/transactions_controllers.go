package controllers

import (
	"net/http"
	"prakerja/configs"
	"prakerja/models"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AdminGetTransactions(c echo.Context) error {
	var transactions []models.Transactions
	result := configs.DB.Find(&transactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Transactions not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: transactions})
	}
}

func AdminGetTransactionsByUserId(c echo.Context) error {
	user_id := c.Param("user_id")
	getTransactions := []models.Transactions{}
	result := configs.DB.Where("user_id = ?", user_id).Find(&getTransactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Transaction not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: getTransactions})
	}
}

func AdminGetTransactionsByTransactionId(c echo.Context) error {
	transaction_id := c.Param("transaction_id")
	getTransactions := []models.Transactions{}
	result := configs.DB.Where("id = ?", transaction_id).Find(&getTransactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Transaction not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: getTransactions})
	}
}

func UserGetTransactions(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YoUr-SeCreT-KeY"), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	var transactions []models.Transactions
	result := configs.DB.Where("user_id = ?", userID).Find(&transactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Transaction not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: transactions})
	}
}
func UserGetTransactionByTransactionId(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YoUr-SeCreT-KeY"), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	var transactions []models.Transactions
	transactionID := c.Param("transaction_id")
	result := configs.DB.Where("id = ? AND user_id", transactionID, userID).Find(&transactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Transaction not found", Data: nil})
	} else {
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: transactions})
	}
}

func UserCreateTransaction(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YoUr-SeCreT-KeY"), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	var transaction models.Transactions
	var incomeCategory models.IncomeCategories
	var expenseCategory models.ExpenseCategories
	var user models.Users
	c.Bind(&transaction)
	if transaction.Type == "income" {
		// Check if income category exists in the database
		err := configs.DB.Where("category = ? AND user_id = ?", strings.ToLower(transaction.Category), userID).First(&incomeCategory).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Invalid income category", Data: nil})
		}

		// Check if balance is sufficient
		err = configs.DB.Where("id = ?", userID).First(&user).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "User not found", Data: nil})
		}

		// Proceed with the transaction
		transaction.Date = time.Now().Format("02-01-2006")
		userIDInt, _ := strconv.Atoi(userID)
		transaction.User_ID = userIDInt
		user.Balance += transaction.Amount

		// Save the updated user balance to the database
		err = configs.DB.Save(&user).Error
		if err != nil {
			// Handle error if failed to save user balance
			return c.JSON(http.StatusInternalServerError, "Failed to update user balance")
		}
		err = configs.DB.Save(&transaction).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to save transaction")
		}

		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Transaction successful", Data: transaction})
	}
	if transaction.Type == "expense" {
		// Check if expense category exists in the database
		err := configs.DB.Where("category = ?", strings.ToLower(transaction.Category)).First(&expenseCategory).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Invalid expense category", Data: nil})
		}

		// Check if balance is sufficient
		err = configs.DB.Where("id = ?", userID).First(&user).Error
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "User not found", Data: nil})
		}
		if transaction.Amount > user.Balance {
			return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Insufficient balance", Data: nil})
		}

		// Proceed with the transaction
		transaction.Date = time.Now().Format("02-01-2006")
		userIDInt, _ := strconv.Atoi(userID)
		transaction.User_ID = userIDInt
		user.Balance -= transaction.Amount

		// Save the updated user balance to the database
		err = configs.DB.Save(&user).Error
		if err != nil {
			// Handle error if failed to save user balance
			return c.JSON(http.StatusInternalServerError, "Failed to update user balance")
		}
		err = configs.DB.Save(&transaction).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to save transaction")
		}
		return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Transaction successful", Data: transaction})
	}

	return c.JSON(http.StatusBadRequest, "Invalid transaction type")
}

func UserDeleteTransactionByTransactionId(c echo.Context) error {
	// Get JWT token from request header
	authToken := c.Request().Header.Get("Authorization")
	splitString := strings.Split(authToken, "Bearer ")
	tokenString := splitString[1]

	// Parse JWT token
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("YoUr-SeCreT-KeY"), nil
	})
	claims, _ := token.Claims.(*models.Claims)

	userID := strconv.Itoa(claims.ID)

	// Get transaction ID from path parameter
	transactionID := c.Param("transaction_id")

	// Check if the transaction exists and belongs to the user
	var transaction models.Transactions
	err := configs.DB.Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, "Transaction not found")
	}

	// Adjust user balance if it is an expense transaction
	if transaction.Type == "expense" {
		// Retrieve user data
		var user models.Users
		err = configs.DB.Where("id = ?", userID).First(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to retrieve user data")
		}

		// Add the transaction amount back to the user balance
		user.Balance += transaction.Amount

		// Save the updated user balance to the database
		err = configs.DB.Save(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to update user balance")
		}
	}
	// Adjust user balance if it is an income transaction
	if transaction.Type == "income" {
		// Retrieve user data
		var user models.Users
		err = configs.DB.Where("id = ?", userID).First(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to retrieve user data")
		}

		user.Balance -= transaction.Amount

		// Save the updated user balance to the database
		err = configs.DB.Save(&user).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to update user balance")
		}
	}

	// Delete the transaction from the database
	err = configs.DB.Delete(&transaction).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete transaction")
	}

	return c.JSON(http.StatusOK, "Transaction deleted successfully")
}
