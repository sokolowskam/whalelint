# Study

Thank you for agreeing to participate in our study. The goal of this stud is to test the usability of our Dockerfile linter. This study consists of 3 parts:
* Dockerfile correction (without a linter)
* Dockerfile correction with assistance of a linter
* Post-study survey

## Dockerfile correction (without a linter)
The Dockerfile is in the directory study/part_1. Your task is to correct or point out all places in code, which you consider to contain a mistake. If the correction involves only changes in the Dockerfile, we encourage you to correct it. If the fix includes creation of additional files, changes in code, moving the execution from build to runtime, write a comment above the line containing a mistake. After you are finished, copy your modified Dockerfile to directory study/part_2. Refrain from modifying the Dockerfile in study/part_1 at later stages of the study.

## Dockerfile correction with assistance of a linter
In this part you will further refine your Dockerfile using our Dockerfile linter. Go to File > Settings > Plugins > Installed, find WhaleLint, click on the checkbox in order to enable the plugin and click Apply. The plugin will annotate your Dockerfile. Try to fix as many problems pointed out by the plugin as possible using its suggestions. Like in the previous stage, if the correction involves only changes in the Dockerfile, we encourage you to correct it. If the fix includes creation of additional files, changes in code, moving the execution from build to runtime, write a comment above the line containing a mistake.

## Post-study survey
In the last stage you will fill out a questionnaire asking questions about the linting process. Please answer them truthfully.