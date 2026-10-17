
# Action Items

## Bugs

+ [x] Transition follow by transition at end of script treats the second transition as action, see sample-05.fountain, noticed in HTML output
    + A: OK, the line begins with "!" so it should be "Action"
+ [ ] sample-05.fountain has the character "DOG" line identified as "dialogue"
+ [ ] sample-07.fountain (charade) has parenthiticals sprinkled in the dialogue, need to support this.  It meakes some dialog show up as Action in the JSON stream.


## Next (v0.1.0)

+ [ ] Check whether the boneyard and scene-number changes (added for the knowledge module's story support, 2026-10-17) have parallels in the other screenplay-format modules and in **scripttool**, which wraps them:
    + **fdx**, **osf** and **scripttool**: do their Fountain conversions drop, keep or mangle a `/* ... */` boneyard now that **fountain** types it as `BoneyardType`? Do they carry `Element.SceneNumber` (FDX `Number` attribute on a scene heading, OSF scene numbers, FadeIn) into and out of Fountain's `#12A#`?
    + Each of them pins an older **fountain** (scripttool v1.0.1, fdx v1.0.1); run their tests against the new release before bumping.
    + Checked 2026-10-17 with **scripttool** built against the local **fountain** 1.0.3: the existing tests of **fdx**, **osf** and **scripttool** pass with it. A Fountain → FDX → Fountain round trip keeps a boneyard, one line or several, as plain text (in FDX and OSF it becomes a visible paragraph, as it did when it was read as Action), and keeps a scene number only inside the heading text (`... #12A#`). Open: **fdx** neither writes the scene's FDX `Number` attribute from `Element.SceneNumber` nor reads it back; look for the same in **osf**; decide whether a boneyard should become a hidden/comment element in FDX and OSF instead of a visible paragraph.
+ [ ] **fountain2html**: boneyard text is written into the page as `<div class="boneyard">`, so cut material is visible; decide whether it should be hidden by default, as Notes and Synopses are

+ [ ] Add test to check to validate parse structure of each `testdata/*.fountain` file
+ [ ] Add demo of using **fountain2json** for creating a script report
+ [ ] Improve **fountainfmt** pretty print options
+ [ ] handle unlabeled title pages
+ [x] Write and **fountain2html**
+ [ ] Add CSS dump option for **fountain2html** 
+ [ ] Add option to render [scrippets](https://fountain.io/scrippets) compatible HTML and CSS


## Someday, Maybe

+ [ ] Add support for Markdown front matter for additiona metadata processing
+ [ ] handle general text (outside of notes, boneyard)
    + this could be handled like front matter in Markdown
+ [ ] definable heading prefixes
+ [ ] reports on screenplay 
+ [ ] fountain2pdf

## Completed

+ [x] Write Parse(src []byte) returning a Fountain Struct and error
+ [x] word wrap for dialogue

