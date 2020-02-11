Overview of Fanning out and Fanning in

- Similar to map reduce

- Fanning out:
    - Takes a job
    - Splits the job into multiple tasks
    - different workers execute the job
    - Write the output into different channels

- Fanning in:
    - Takes the intermediate channels
    - Combines and writes into one channel 