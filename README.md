Galaxy Merchant Trading Guide
===========================

This is repository of the Galaxy Merchant Trading Guide problem solution created by [Naufal Farras][1]. The solution of problem is typed in [Go Programming Language][2].

# Requirement
The solution requires Go >= 1.8.

# Problem Statement
You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot).Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you.The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate translation between them.Roman numerals are based on seven symbols.

#### Symbol Value
* I 1
* V 5
* X 10
* L 250
* C 100
* D 500
* M 1,000

Numbers are formed by combining symbols together and adding the values. For example, MMVI is 1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the largest values. When smaller values precede larger values, the smaller values are subtracted from the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.) "D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of 1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately. Inthe above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

[Source: Wikipedia Roman Numbers][3]

Input to your program consists of lines of text detailing your notes on the conversion between intergalactic units and roman numerals. **You are expected to handle invalid queries appropriately**.

### Test input:
> glob is I
> prok is V
> pish is X
> tegj is L
> glob glob Silver is 34 Credits
> glob prok Gold is 57800 Credits
> pish pish Iron is 3910 Credits
> how much is pish tegj glob glob ?
> how many Credits is glob prok Silver ?
> how many Credits is glob prok Gold ?
> how many Credits is glob prok Iron ?
> how much wood could a woodchuck chuck if a woodchuck could chuck wood ?

### Test Output:
> pish tegj glob glob is 42
> glob prok Silver is 68 Credits
> glob prok Gold is 57800 Credits
> glob prok Iron is 782 Credits
> I have no idea what you are talking about

# Solution Approach
The flowchart of solution looks like below
![Solution Flowchart][4]

Input is divided into 4 types (categorized by using Regex):
### Type 1
Input categorized into type 1 looks like *glob is I* with regex
> "[a-zA-Z]+ is [IVXLCDMivxlcdm]{1}"
The type of input tells convention between intergalactic numerals and roman numerals. For example
> glob is I means glob equals to I
> prok is V means prok equals to V
> etc
The name of function for the task of type is *interRomanInput()*.

### Type 2
Input categorized into type 2 looks like *glob glob Silver is 34 Credits* with regex
> "[a-zA-Z ]+ is \\d+ [cC]redits"
The type of input tells credits value of the metal. For example
> glob glob Silver is 34 Credits means 2 (glob glob is II as following convention) silver equals 34 Credits
> glob prok Gold is 57800 Credits means 4 (glob prok is IV as following convention) gold equals 57800 Credits
> etc
The name of function for the task of type is *metalsCreditValueInput()*.

### Type 3
Input categorized into type 3 looks like *how much is pish tegj glob glob ?* with regex
> "how much is [a-zA-Z ]+ \\?"
The type of input is a question of intergalactic numerals value (through roman numeral). It will return answer the value of intergalactic numerals asked. For example
> how much is pish tegj glob glob? will return pish tegj glob glob is 42
> etc
The name of function for the task of type is *interRomanQuestionInput()*.

### Type 4
Input categorized into type 4 looks like *how many Credits is glob prok Silver ?* with regex
> "how many [cC]redits is [a-zA-Z ]+ \\?"
The type of input is a question of some of metals credits value. It will return answer the value of some of metals. For example
> how many Credits is glob prok Silver ? will return glob prok Silver is 68 Credits
> how many Credits is glob prok Gold ? will return glob prok Gold is 57800 Credits
> etc
The name of function for the task of type is *interCreditsValueQuestionInput()*.

If input can't be categorized into the following types, program will return "I have no idea what you are talking about".

# Tests
The tests are done manually not programmatically. The tests are divided into
## Unit Tests
#### RomanIntoNumberConv() in './romans/romans.go'
Convert roman digits into numbers (int)
| Test Case | Expected |  Actual |
| --------- | -------- | ------- |
| VIII      |    8     |    8    |
| IX        |    9     |    9    |
| VX        | *Error*  | *Error* |
| IIX       | *Error*  | *Error* |
| IM        | *Error*  | *Error* |
| LCM       | *Error*  | *Error* |

#### intergalacticIntoRoman() 
Convert intergalactic digits into roman digits
**Known glob is I, prok is V, pish is X, tegj is L**
|    Test Case   | Expected |  Actual |
| -------------- | -------- | ------- |
| glob glob      |   II     |   II    |
| pish glob prok |   XIV    |   XIV   |
| tegj           |   L      |   L     |
| mush pish      | *Error*  | *Error* |

#### inputClassification()
Classify input into 4 types
|                                Test Case                                | Expected |  Actual |
| ----------------------------------------------------------------------- | -------- | ------- |
| glob is I                                                               |  Type 1  |  Type 1 |
| glob glob Silver is 34 Credits                                          |  Type 2  |  Type 2 |
| how much is pish tegj glob glob ?                                       |  Type 3  |  Type 3 |
| how many Credits is glob prok Iron ?                                    |  Type 4  |  Type 4 |
| how much wood could a woodchuck chuck if a woodchuck could chuck wood ? | *Error*  | *Error* |
| aaaaaaaegegwebrwhwr                                                     | *Error*  | *Error* |

# Installation
#### 1. Install NodeJS
> **Note:**
> Before you install, please check first by run a command
> ```
> node --version
> ```
> If the version of node >=8.5.0, you can **skip** this step.

1. If you are Windows user, please download [here][3]. Choose the right Windows version (32/64 bit) based on your computer. Please check [here][4] to get the version of Windows.
2. If you're Ubuntu user, run command

    ```
    curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
    sudo apt-get install -y nodejs
    ```
3. For another OS, you can check [here][5].

#### 2. Clone The Project
> **Note:**
> You must install [**git command**][6] first before do the step.

You can store the project file wherever you want. Clone the project by run a command
```
git clone https://gitlab.com/naufalfmm/server_side.git
```
Before you clone the project, you have to be in the directory you want to store the project file.

#### 3. Install The Dependencies
> **Note:**
> You have to be in the directory of the project file stored.

Install the dependecies by run a command
```
npm install
```
If you use yarn, please run command below.
```
yarn
```

#### 4. Run The App!
The application is ready to serve on 7777. Please type and run command below to start the application
```
node app.js
```
or
```
npm start
```

# API Documentation
## **Admin**
---
### **Admin Login**
  Login for admin role. Returns token for "auth" header using JWT.
  ```
  POST /api/admin/login
  ```
#### Request Body:
  ```json
  {
      "username":"",
      "password":""
  }
  ```
#### **Success Response:**
* **Code:** 200 <br />
  **Content:** 
  ```json
  {"status": true, "token": "", "message": "Happy Login"}
  ```
#### **Error Responses:**
* **Code:** 500 Internal Server Error <br />
  **Content:** 
   ```json
   {"status": false, "message": "Internal Server Error"}
   ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
---
## **Node**
---
### **Node Login**
  Login for node role. Returns token for "auth" header using JWT.
  ```
  `POST` /api/device/login
  ```
#### **Request Body:**
  ```json
  {
      "username":"",
  }
  ```
#### **Success Response:**
* **Code:** 200 <br />
  **Content:** 
  ```json
  {"status": true, "token": "", "message": "Happy Login"}
  ```
#### **Error Responses:**
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
### **Add/Register Node**
  Register new node for a lane. Each lane is just for one node. Node **only** can be registered by Admin.
  ```
  POST /api/device/add/{laneId}
  ```
#### Path Parameters
| Parameter name | Value   | Description | Additional |
| -------------- | ------- | ----------- | ---------- |
| laneId         | integer | Lane id     | Required   |

#### **Success Response:**
* **Code:** 201 <br />
  **Content:** 
  ```json
  {"status": true,  "message": "Device Successfully Created"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 409 Conflict <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unique Constraint Error"}
  ```
  or
  ```json
  {"status": false, "message": "Lane Id Error"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
### **Delete Node**
  Delete node registered by admin.
  ```
  DELETE /api/device/delete/{deviceId}
  ```
#### Path Parameters
| Parameter name | Value   | Description | Additional |
| -------------- | ------- | ----------- | ---------- |
| deviceId       | integer | Device id   | Required   |

#### **Success Response:**
* **Code:** 200 <br />
  **Content:** 
  ```json
  {"status": true,  "message": "Device is Successfully Deleted"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 403 Forbidden <br />
  **Content:** 
  ```json
  {"status": false, "message": "Device Doesn't Exist"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
### **Create Node Config**
  Create config file of node.
  ```
  POST /api/device/create/{deviceId}
  ```
#### Path Parameters
| Parameter name | Value   | Description | Additional |
| -------------- | ------- | ----------- | ---------- |
| deviceId       | integer | Device id   | Required   |

#### **Success Response:**
* **Code:** 201 <br />
  **Content:** 
  ```json
  {"status": true,  "message": "Config File Created Successfully"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 403 Forbidden <br />
  **Content:** 
  ```json
  {"status": false, "message": "Device Doesn't Exist"}
  ```
* **Code:** 404 Not Found <br />
  **Content:** 
  ```json
  {"status": false, "message": "Device Not Found"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
### **Download Node Config File**
  Download node config file created by server.
  ```
  GET /api/device/download/{deviceId}
  ```
#### Path Parameters
| Parameter name | Value   | Description | Additional |
| -------------- | ------- | ----------- | ---------- |
| deviceId       | integer | Device id   | Required   |

#### **Success Response:**
* **Code:** 200 <br />
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 403 Forbidden <br />
  **Content:** 
  ```json
  {"status": false, "message": "Device Doesn't Exist"}
  ```
* **Code:** 404 Not Found <br />
  **Content:** 
  ```json
  {"status": false, "message": "File Not Found"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
---
## **Road**
---
### **Add Road**
  Add new road to *database*.
  ```
  `POST` /api/road/add
  ```
#### **Request Body:**
  ```json
  {
      "name":"",
      "lanes":,
      "paths":
  }
  ```
#### **Success Response:**
* **Code:** 201 <br />
  **Content:** 
  ```json
  {"status": true, "message": "Road Successfully Added"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
---
## **Lane**
---
### **Add Lane**
  Add new lane of road to *database*.
  ```
  `POST` /api/lane/add
  ```
#### **Request Body:**
  ```json
  {
    "th":,
    "width":,
    "cap":,
    "path":,
    "road_id":
}
  ```
#### **Success Response:**
* **Code:** 201 <br />
  **Content:** 
  ```json
  {"status": true, "message": "Lane Successfully Added"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
  or
  ```json
  {"status": false, "message": "Bad Request"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```
---
## **Traffic Data**
---
### **Post traffic data**
  Post traffic data to the *database*.
  ```
  `POST` /api/data/post
  ```
#### **Request Body:**
  ```json
  {
    "data": "encrypted_data"
  }
  ```
#### **Success Response:**
* **Code:** 201 <br />
  **Content:** 
  ```json
  {"status": true, "message": "Data is Successfully Added"}
  ```
#### **Error Responses:**
* **Code:** 400 Bad Request <br />
  **Content:** 
  ```json
  {"status": false, "message": "Data Incomplete"}
  ```
  or
  ```json
  {"status": false, "message": "Bad Request"}
  ```
* **Code:** 401 Unauthorized <br />
  **Content:** 
  ```json
  {"status": false, "message": "Unauthorized"}
  ```
* **Code:** 403 Forbidden <br />
  **Content:** 
  ```json
  {"status": false, "message": "Device Doesn't Exist"}
  ```
* **Code:** 500 Internal Server Error <br />
  **Content:** 
  ```json
  {"status": false, "message": "Internal Server Error"}
  ```

[1]: https://gitlab.com/naufalfmm
[2]: https://golang.org/
[3]: http://en.wikipedia.org/wiki/Roman_numerals
[4]: https://github.com/naufalfmm/galaxy_merchant_trading_guide/blob/master/docs/solution_flow.png