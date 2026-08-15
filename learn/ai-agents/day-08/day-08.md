# What a multi-agent system is
A: multi-agent ststem is each agent have own role, tools, and prompt. And they can complete their tasks

# Each agent's role + tools
A: I have 2 agent: 

- barista agent: this agent use menu price and calcualte order total tools. It can complete RotaCoffee barista tasks.
- researcher agent: this agent use DuckDuckGoSearchRun tools. It can search query in web.

# How Routing works
A: Routing has some keywords. If these keywords present in the prompt, it is routed to the barista agent. Otherwise, its routed to the researcher agent.


# Test Result:

[barista] A latte is 45 TL.
[barista] The total for 2 lattes and 1 espresso is 125 TL.
[researcher] Espresso originated in **Italy** in the late 19th and early 20th centuries. 

Here is a brief history of how it came to be:

* **1884:** Angelo Moriondo of Turin, Italy, patented the first steam-driven espresso machine designed to brew coffee quickly for customers.
* **1901–1905:** Luigi Bezzera improved the design by adding portafilters and multiple brew heads. Desiderio Pavoni bought Bezzera's patents and commercialized the first espresso machine, *La Pavoni*, introducing it to the world at the 1906 Milan Fair.
* **Meaning:** The term *"espresso"* comes from the Italian word for "pressed out" or "express," referring to how coffee is forced through finely-ground beans under high pressure in a short amount of time.

If you'd like to check the price of an espresso or any other drink on our menu, feel free to ask!


[researcher] While coffee with milk (*caffè latte*) has been enjoyed in European homes for centuries, the modern espresso-based **Caffè Latte** as served in coffee shops today was popularized in the 1950s by **Lino Meiorin**, a co-owner of the **Caffe Mediterraneum** in Berkeley, California.

Italian-born Meiorin adapted the traditional Italian cappuccino—which many American customers found too strong—by adding extra steamed milk, eventually putting "Caffè Latte" on the menu.

If you'd like to check our menu price for a latte or place an order, let me know!