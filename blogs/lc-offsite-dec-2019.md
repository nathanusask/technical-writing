# Lucky Charms 2020 Kick-off offsite

## Table of Contents
* [Overview](#overview)
* [Ice-breaker](#ice-breaker)
* [Company 3-year strategy and goals and Marketplace divisional goals](#company-3-year-strategy-and-goals-and-marketplace-divisional-goals)
* [Team principles](#team-principles)
    * [Continuous value delivery](#continuous-value-delivery)
    * [System of record](#system-of-record)
    * [API first](#api-first)
    * [Deep integration](#deep-integration)
* [Workshops](#workshops)
    * [Deep integration vs friction-free UX](#deep-integration-vs-friction-free-ux)
    * [Vendor struggles](#vendor-struggles)
    * [Partner struggles](#partner-struggles)

## Overview
The offsite took place on Thursday, December 19th 2019 at the Living Sky Cafe on the 3rd Ave S.
We had four developers of the team, sales engineer Jared Eason, product manager Levi Johnson,
engineering manager Hemant Naidu, UX designers Cam Johnson and Brendan Swalm for almost the entire time of the day.
We also invited guests from other divisions to join us. Special thanks go to Ed O'Keefe, Nolan Cline, Sarah Shell,
Chiara Traversa, and Brayden Fehr for their external perspectives.
We've made available the [agenda](https://docs.google.com/document/d/1rA8XXQsur9CP0pRIyyVudcRAeVJIcheSFOXiAtG-DAI/edit)
and [notes](https://docs.google.com/document/d/194ooCx13y4LL-bi_6yXcnw-oZ11NWrwBLiQqj5_Y3GU/edit#heading=h.52rkah2hanpn)
for anyone who is interested in the details.

## Ice-breaker
To kick off the day, Corey planned a fun game for everyone and the game was to guess who wrote what.
Briefly, this game asked each participant to answer the questions in their assigned cards and then add
their own names along with a different one from one of the other attendants. For instance, Nathan was asked who
would he invite to host a late night show; Nathan answered Conan O'Brien (yeah you can't deny Conan's self-deprecation comedy
can you?); Nathan wrote down his name behind Levi who Nathan thought would give the same answer; then people were asked
who wrote down the answer from the two candidates; i.e., Nathan or Levi; sadly more people chose Levi over Nathan lol.
It was a great 15-minute game that brought everyone closer.

## Company 3-year strategy and goals and Marketplace divisional goals
For this section, Levi did a quick recap of the company's 3-year strategies and goals, by highlighting what we wanted to accomplish
separately by the end of each year. Among the various goals, the one that specifically focuses on the Marketplace division
was to build Marketplace infrastructure to support new categories. To measure the success of these goals, Ed added
Gross Merchandise Volume (GMV) and North Star Metric. Two strategic pillars were laid out after: capture GMV and
allow SMBs to purchase direct from Cloud Brokers via our platform.
To capture GMV we would want to monetize transactions for quarterly revenue and profit.
To allow SMBs to purchase direct from Cloud Brokers via our platform we would want to do easy integration for independent software vendors (ISV).
The ISV integration was actually what the Lucky Charms team were going to do for 2020. Upon this offsite, Levi listed out four global brands that
the team would be integrating in 2020 and they are Google Cloud (Google version of the Amazon Web Service), FreshBooks (an accounting solution),
Symantec (cybersecurity software and services), and Deputy (a scheduling software & time clock app). For Q1, the team
would target on FreshBooks integration first.
(Phew~~I hope what I've written so far for this section makes sense...if not or you wanna find more details,
please refer to the [presentation](https://docs.google.com/presentation/d/13HEVyUkJl_VOy2dvRZn60fX1yila8BIn3Gz5tEn63v4/edit#slide=id.g7ba249d95d_0_1).)

## Team principles
This session was more of a retrospective of the past year on what we did well and what we didn't.
During this session, attendants were asked to write down items for each category, sort the items, and dot vote what they
think were the most important. Finally, we picked the top 2 from each category and formed new team principles for 2020.
The top 2 good things we did in 2019 were continuous value delivery and being the system of record; on the other hand,
the top two that we didn't do well were API and deep integration.

### Continuous value delivery
From Jared's perspective, the products we delivered in the past were what attracted our customers; the customers came to
us for values; they acknowledged that everything in our platform was accessible and they had the faith that they
would grow with us. With this positive feedback in mind, we would want to keep delivering values to our suppliers;
i.e., partner vendors and cloud brokers.

### System of record
In the past, we kept records of all orders; e.g., activation and deactivation. In addition, we made available the display
of these records as well as other critical information such as active status. In 2020, we would want to integrate data to
to improve data-driven user experience so that we make our platform the system of engagement as well. 

### API first
During the past, our main focus was integrating products into our platform using their APIs,
not aware of the possibility that other companies would want to integrate with us in future
so we didn't pay attention to providing APIs for potential partners. In light of this issue,
we agreed on being consistent with using API for strategy to design.

### Deep integration
In the past, our integration work mainly revolved around a limited set of operations;
i.e., account create/delete, SSO, and billing. However, these operations might have not provided
highly smooth user experience and some might even find it difficult to navigate between our platform
and the product. The main reason we struggled on this item was mainly because the lack of APIs from
the brand we integrated or the brand was difficult to work with; e.g., we found Microsoft's support team
was not cooperative in providing us actionable solutions. As much as we wanted to improve our integration work,
it wasn't easy to come up with a tangible goal to mitigate this issue at that moment.

## Workshops
Three workshops took place during the day; one before lunch and the other two in the afternoon.
Parker facilitated the one-and-half-hour workshop to discuss pros and cons of doing deep integration versus friction-free UX.
In the afternoon, our guests shared their thoughts in hig-level to discuss the struggles they saw from partners and vendors
in the past year.

### Deep integration vs friction-free UX
Deep integration means data are synchronized between our platform and the global brands platforms.
Also, it means that the user portal we provide for partners has SSO that allows them to access their accounts on
the other end; i.e., the brand itself. Such a user portal is meant to provide navigation for users to go back and forth
between our platform and the brand. However, our integration work didn't go with that depth and somehow partners might
have constantly lost track during their use of our platform.

Jared mainly focused on what was not working and what did not work properly.
He acknowledged that not all of them were related to the team; various causes could make things go wrong.
Fro instance, sometimes the issues were not taken care of on the brands side such as the unmatched G Suite license number issue;
or it could be the things we did were the those our cloud brokers did not care about the most.
However, there were issues that were our fault. Products such as G Suite and MS-Office had confusing setups.
Specifically, majority of our cloud brokers were marketplace agents who did not understand our platform very well.
For example, we had these portals, but the cloud brokers did not fully understand the workflow; e.g., our G Suite portal
could do some of the admin stuff but not all of it, which became a UX issue.
Things such as monthly billing and transfer were not well integrated either; e.g., G Suite was not integrated with monthly billing
and MS-office currently did not have transfer ready.
Moreover, ee still had missing features for products such as Office subscriptions and we were missing opportunities
because of this incomplete integration. 
Further, the SSO might not be a good user experience, especially for G Suite and Office,
because the way email was viewed by the cloud brokers was different than going to google directly which was more intuitive
for them; in other words, they might want to have fixed credentials instead of reading emails all the time to
get temporary access. Sometimes they would guess why things were not working was because of the poor SSO.
Finally, Jared suggested that the way we integrated products should be more proactive,
because we didn't want to miss opportunities.
One way to accomplish this, as he pointed out, was to look at our competitors and dig into their features;
in other words, we should consider things before cloud brokers complain.

Parker then asked attendants to write down the pros and cons for doing deep integration and then dot vote
what were the most important ones. After a while, we picked the following tops:
For pros, people agreed on that deep integration would make our platform a one-stop shop and hence would
bring in more opportunities. In addition, people agreed on that deep integration would definitely add more value to the company.
For the cons, people agreed on that deep integration would end up with features not being used.
However, Levi suggested that it would be rare and that we do more competitor research.
Also, time and MVP could be another drawback for making deep integration happen.

### Vendor struggles

### Partner struggles
