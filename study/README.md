# Study

Thank you for agreeing to participate in our study. The goal of this study is to test the usability of our Dockerfile linter. This study consists of 3 parts:
* Dockerfile correction (without a linter)
* Dockerfile correction with assistance of a linter
* Post-study survey

## Dockerfile correction (without a linter)
The Dockerfile is in the directory [study/part_1](part_1/Dockerfile). Your task is to correct or point out all places in code, 
which you consider to contain a mistake. If the correction involves only changes in the Dockerfile, 
we encourage you to correct it. If the fix includes creation of additional files, changes in code, 
moving the execution from build to runtime, write a comment above the line containing a mistake. 
After you are finished, copy your modified Dockerfile to [study/part_2/Dockerfile](part_2/Dockerfile). For the next step 
you will be working with this copied Dockerfile.
Refrain from modifying the Dockerfile in [study/part_1](part_1/Dockerfile) at later stages of the study.

## Dockerfile correction with assistance of a linter
In this part you will further refine your Dockerfile (copied in location [study/part_2](part_2/Dockerfile)) using our Dockerfile linter. 
Go to `File > Settings > Plugins > Installed`, find `WhaleLint`, click on the checkbox 
in order to enable the plugin and click `Apply`. 
The plugin will annotate your Dockerfile. Try to fix as many problems pointed out 
by the plugin as possible using its suggestions. Like in the previous stage, 
if the correction involves only changes in the Dockerfile, we encourage you to correct it. 
If the fix includes creation of additional files, changes in code, moving the execution 
from build to runtime, write a comment above the line containing a mistake.

## Post-study survey
In the last stage you will fill out a questionnaire asking questions about the linting process. 
You will be answering 16 questions for which you have to rate each question
on a scale between Strongly Agree to Strongly Disagree.
Please answer them truthfully.
Click [here](https://forms.gle/cvxQBxrBgKzNs4JV9) for the post-study survey.