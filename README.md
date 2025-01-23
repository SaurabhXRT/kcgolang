# kcasssignment 

## Description
IN This project Users can submit a job containing multiple image URLs, and the system will process these images asynchronously. Each job's status (e.g., `ongoing`, `completed`, `failed`) can be queried at any time. The project is designed to handle image processing tasks, calculate image perimeters, and simulate GPU-intensive processing.

---

## Assumptions
- Each submitted job includes a list of visits, and each visit contains a list of image URLs.
- A job is considered `completed` only if all images in the job are processed successfully.
- If any image in a job fails to process, the entire job is marked as `failed` with error including that storeid.
- Image processing takes GPU-intensive tasks with a fixed 30-second delay.
- The image perimeter calculation uses 2D dimensions (height and width) of the image's bounds.
- The system does not persist data; all processing is stored in memory for the duration of runtime.
- I have used vscode 
---
## Installing and Testing Instructions

- use command go mod init kcasssignment and go mod tidy 
- And use go run main.go to run the server if using docker use docker-compose up --build
- There are two endpoint /api/submit and /api/status?jobid=
- once we submit the job it will give response with created jobid and a asynchronous function will start to process the image with timedelay of 30second and meanwhile if check the status of that jobid it will show status ongoing and if failed then it will give response with storeid and error
- I have used this request payload {
   "count":2,
   "visits":[
      {
         "store_id":"S00339218",
         "image_url":[
            "https://www.gstatic.com/webp/gallery/2.jpg",
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "time of store visit"
      },
      {
         "store_id":"S01408764",
         "image_url":[
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "time of store visit"
      }
   ]
}


# if given more time
- I would have used database like postgresql using goorm instead of temporarily storing the jobs in map

