# TASK 1 

## Went well
- Carried one project (RotaCoffee) through the whole SDLC, so each day built on the last instead of feeling random.
- Got comfortable writing user stories and testable acceptance criteria in English.
- Learned to separate the problem from the solution instead of jumping straight to tech.

## Improve
- The early SDLC concepts were confusing at first — I mixed up the phases like design vs implementation and what each one actually produces.
- I sometimes accepted answers before fully understanding them (like the testing task).
- My Git habits are still shaky — case-sensitivity and staging tripped me up.


## Actions
- Write the first draft of each task in English myself, then ask for corrections.
- When something doesn't click, stop and ask for one small example before moving on.
- Practice the basic Git flow (branch → commit → push) a few times until it's automatic.


# TASK 2

Complaint: I mixed up the early SDLC phases (design vs implementation) and what each produces.
Action: Re-read my Day 1–4 notes and write a one-line summary of what each phase produces, in my own words.
Owner: me
By when: end of Day 16

Complaint: My Git staging and case-sensitivity caused failed commits.
Action: Redo the branch → small commits → push flow on a practice branch until it's automatic.
Owner: me
By when: Aug 15


# TASK 3

Metric: Escaped defects — bugs that reach prod because testing missed them.

Why it matters: RotaCoffee handles payments, so a bug reaching real customers means 
real money going wrong. For my learning, it tells me whether my tests actually cover 
the risky paths, not just the happy path.

How I'd track it: count how many bugs are found in prod vs caught before release each 
month, and watch if the prod number goes down as my testing habits improve.


# TASK 4

- A clear alert that names the service and what's failing
- A link to the runbook with rollback steps
- A list of recent deploys so I know what changed
- Error logs with a full stack trace
- A dashboard showing error rate and response times
- Who to escalate to if I'm stuck
- Access and permissions already set up so I'm not locked out mid-incident