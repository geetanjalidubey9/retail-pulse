
# Project Documentation for Retail Pulse Image Processing Service

## *. Project Description

Retail Pulse provides an image processing service to process thousands of images collected from stores. The service calculates the perimeter of each image, simulates GPU processing with a random sleep time between 0.1 to 0.4 seconds, and stores the result for each image. The project includes two main features:

- **Submit Job API** (`/api/submit/`): Accepts a request to process multiple images, each associated with a store, and returns a job ID.
  
- **Get Job Status API** (`/api/status?jobid=`): Allows the user to check the status of the job. The job status can be `completed`, `ongoing`, or `failed`.

This service simulates a job queue system that processes images, calculates a value for each, and can be checked for progress through job status updates.

---

## **2. Assumptions**

- The service receives a valid list of image URLs and store IDs. If a `store_id` or image URL is invalid, the job will fail for that store.
- The perimeter calculation for an image is a simple formula based on the height and width of the image (2 * [Height + Width]).
- The service assumes that the image download and calculation will be successful for most jobs. In case of failure, an error message will be returned for the specific store.
- The random sleep time between image processing is simulated to imitate GPU processing delays.
- The project assumes that the external APIs (for downloading images) are reliable and responsive.
- If any field in the payload is missing or the `count` doesn’t match the length of the `visits` array, the job creation will fail.

## **3. Installing and Testing Instructions**

### **Setting up the project**:

1. **Clone the repository**:

   git clone https://github.com/yourusername/retail-pulse-image-processing.git
   cd retail-pulse-image-processing

2. **Set up Go Modules**:
   Ensure you have Go installed on your system. Initialize Go modules if you haven't already:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   To run the application without Docker, use the following command:
   ```bash
   go run main.go
   ```

4. **Docker Setup**:
   If you prefer to use Docker, make sure you have Docker installed and then run:
   ```bash
   docker build -t retail-pulse-image-processing .
   docker run -p 8080:8080 retail-pulse-image-processing
   ```

### **Testing**:
You can use **Postman** to test the API endpoints:

1. **Submit Job** (`POST /api/submit/`):
   - Send a `POST` request with the payload containing `count` and `visits`.
   - Example:
     ```json
     {
        "count": 2,
        "visits": [
           {
              "store_id": "S00339218",
              "image_url": [
                 "https://www.gstatic.com/webp/gallery/2.jpg",
                 "https://www.gstatic.com/webp/gallery/3.jpg"
              ],
              "visit_time": "2024-11-20T12:00:00Z"
           }
        ]
     }
     ```
   
2. **Get Job Status** (`GET /api/status?jobid=`):
   - Send a `GET` request with the `jobid` received from the response of the `POST /api/submit/` request.
   - Example: `GET http://localhost:8080/api/status?jobid=1`

3. **Expected Responses**:
   - **POST /api/submit/**: Returns a job ID (e.g., `{ "job_id": 1 }`).
   - **GET /api/status**: Returns job status (e.g., `completed`, `ongoing`, or `failed`).

---

## **4. Work Environment**

- **Operating System**: Ubuntu 22.04 LTS (Linux) / Windows 10
- **Text Editor/IDE**: Visual Studio Code / GoLand
- **Libraries and Tools Used**:
  - **Go** (version 1.18 or above)
  - **Docker** (for containerization)
  - **Gin Web Framework** (for HTTP API server)
  - **Go Modules** (for dependency management)
  - **Postman** (for API testing)
- **Dependencies**:
  - `github.com/gin-gonic/gin` (Gin Web Framework)
  - `github.com/go-playground/validator/v10` (for request validation)

---

## **5. Improvements for the Future**

If given more time, the following improvements could be made:

1. **Database Integration**:
   - Integrate a database (e.g., MongoDB or PostgreSQL) to store job details, image processing results, and logs for better tracking and querying.
   
2. **Job Queue Management**:
   - Implement a job queue to handle large volumes of image processing jobs efficiently. Use tools like **RabbitMQ** or **Redis** for managing queues and worker processing.

3. **Error Handling and Retry Logic**:
   - Enhance error handling by adding retry mechanisms for transient issues like network failures or temporary downtime when downloading images.


4. **Authentication and Authorization**:
   - Add user authentication (e.g., using JWT tokens) to secure the API and allow only authorized users to submit jobs.



5. **Image Validation and Preprocessing**:
   - Implement validation to check if the provided image URLs are valid and if the images are of acceptable formats before processing.

6. **Cloud Deployment**:
   - Deploy the service on cloud platforms like AWS, Google Cloud, or Azure to handle larger workloads and scale the service horizontally.


