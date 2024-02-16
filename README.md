# Automated Programming Project

For the project contained within this repository, we put ourselves in the position of a data scientist working for a company whose leadership team is interested in evaluating whether it is over-paying the data scientists, data engineers, and data analysts at the company.  Specifically, the company's leadership points out that recent advancements in natural language processing (NLP) have led to the release of many new tools that can write or edit entire programs for data scientists.  Given that such tools have been streaming onto the market, the company's leadership wonders whether the company still truly needs to be paying $120,000 and $175,000 annual salaries to junior Golang programmers and senior software engineers, respectively.  The company's leadership requests that the data scientist learn more about three newly developed automated programming tools: 1) an Automated Code Generator , 2) an AI-Assisted Programming Tool, and 3) a large language model (LLM) capable of generating code (AI-Generated code). The data scientist must report to the company's leadership about the effectiveness of each of these tools and the degree to which the company may be able to save money on programmer salaries due to these tools automating the work associated with these positions.

For this evaluation, the data scientist focuses on leveraging these three automated programming tools to calculate linear regression models for each of the four Anscombe Quartet datasets.  For the first (Automated Code Generator) phase of the experiment, the data scientist examines the extent to which he's able to leverage the Golang Generate command to write the regression-calculating program.  For the second (AI-Assisted Programming) phase of the experiment, the data scientist evaluates the extent to which the Github Copilot application is able to improve the performance of a previously written Golang program focused on calculating regression models for Anscombe Quartet data (Desilets 2024).  For the third (AI-Generated Code) phase of the experiment, the data scientist requests that ChatGPT write the entire program for him and saves the original text of the conversation with the LLM in this repository as a separate text file. For each of these three automated programming methods, the data scientist also performs 100 iterations of the regression-calculating experiments and records the runtime for each experimental trial.  The data scientist then prints the total and average runtime for these 100 experimental trials to an output text file - each of which is included within this repository.

After leveraging Golang Generate for the first (Automated Code Generator) evaluation,  the data scientist observed the strengths and weaknesses of this command for reducing the programming workload of data scientists.  The data scientist created the "AutomatedCodeGenerationCode.go" file within this repository from scratch using Go Generate.  Interestingly, as the data scientist wrote the program, Go Generate would predict each line of code that he was trying to write in the same way that Word and Gmail autocomplete features aim to predict users' sentences.  In fact, if the data scientist typed comments explaining what he wanted the subsequent lines of code to complete, then Go Generate could accurately predict entire blocks of code he wanted to write. Furthermore, this command seemed intelligent enough to incorporate outside information into its predictions.  (Specifically, Go Generate generated code defining all the data points in the Anscombe Quartet just because the data scientist indicated in a comment that he wanted to work with Anscombe Quartet data). Despite these successes, Go Generate did still have some slight weaknesses.  Specifically, Go Generate's predictions were not always fully accurate and Go Generate still did rely on the data scientist to oversee the overall structure of the code quite a bit as the file was being developed.  In the end, this program, correctly calculated the regression lines as y = 0.5*x + 3, and completed these calculations with an average experimental trial runtime of 2,073 microseconds - the slowest runtime key performance indicator (KPI) resulting from our three evaluations.  Though the Go Generate command may not have yielded the strongest results computationally or have written the program entirely independently, this program was probably the second most successful methodology in reducing the overall programming workload for the data scientist.

[INSERT MODEL 2 RESULTS]

[INSERT MODEL 3 RESULTS]

[INSERT MODEL 4 RESULTS]


References

Desilets, Steve. 2024. "Benchmarking Project / GoBenchmarkingCode.go". Github. https://github.com/Steve-Desilets/Benchmarking-Project/blob/main/GoBenchmarkingCode.go 


The Batch file within this repository can be called by the Windows Command Prompt Application to call the Python, R, and Go code (also contained within this repository) to execute the benchmark testing experiment described above.  We see that Go, Python, and R were able to successfully calculate the correct values for the summary statistics (including minimum, maximum and mean) for each of the seven variables in the CSV input dataset.  Each of the three programming languages also conducted 100 iterations of these calculations, so that we could gain a representative sample of benchmark test trial times for each programming language.  The three programs printed their summary statistics calculations, as well as their eperimental runtimes to the output text files contained within this repository.

By examining the results, we see that the average benchmarking experiment trial runtimes for Go, Python, and R were 23,146 microseconds , 27,832 microseconds, and 100,810 microseconds, respectively.

If I were to make a recommendation to management based on the findings of this study, I would encourage leadership to consider its organizational objectives when deciding whether to leverage Python, R, or Go for future data science projects.  That's because each of these three languages comes with its own distinct advantages and disadvantages.  For example, if management were purely focused on the efficiency of a program when running code to calculate summary statistics from the Command Prompt Application, then I would recommend leveraging R, since it performed the strongest in the benchmark experiment described above.  However, leadership may still want to consider other benefits of leveraging each language.  For example, Go enforces coding best practices that cause programmers to write code in an idiomatic and readable fashion.  Also, Python creates data visualizations that are the most appealing of these three languages.  Meanwhile, R is perhaps the most commonly used language for data science, so there exist many online resources to help programmers write and debug R code, and the firm might have relatively little difficulty recruiting talented R programmers to the company.  In these ways, each of the three languages examined in this study have their own strengths and weaknesses.  Still, if the company is purely focused on the results of this Command Line Application benchmarking study that examined the ability of each language to import CSV files, calculate summary statistics, and generate output text files, then the company should elect to use R for its data science operations, since R was the most computationally efficient language.
