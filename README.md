# Blume-hack

#Objective: 

    Build a platform to extract item & store details from billed receipt
    
 #Tech stack
       1. AWS API Gateway
       2. Golang
       3. AWS Lambda
       4. AWS S3
       5. AWS Rekognition
       6. Kafka
       7. Elasticsearch
       8. Kubernetes
    
    
 #Cloud Providers
       1. GPC
       2. AWS   
       
      P.S: We ran Kafka & Elasticsearch on GCP. AWS only provide 1 Core, 1 RAM for free tier :(
  
 # Design Diagram
 
 ![Alt text](design.png?raw=true "Title") 
 
 #Future Roadmap
 
 1. Improve data normalisation from cluttered image to text
 2. Rotation check of the images.
 3. Improve quality of image & then extract the text  data
 4. Improve noise from retrieved textual data
 