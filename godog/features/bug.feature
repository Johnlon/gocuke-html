Feature: User Registration
  In order to enhance communication with colleagues
  As an employee
  I want to register for the instant messaging (IM) software.

  Rule: Valid Registration
  In order to successfully register
  As a user
  I need to provide a valid mobile number and a valid verification code and a valid password

    Scenario: should register successfully.
      Given registration information:
        | mobile number | verification code | password |
        | 18827073676   | 1111              | 123      |
      When I send "POST" request to "/sso/api/register/basic"
      And the response code should be 200
      And the response should match json:
      """
      {
          "code": 0,
          "msg": "成功",
          "data": ""
      }
      """
