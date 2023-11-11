# options: 
#  -c to continue after failed test
#  -d to execute command in debug mode

export PATH=$PATH:.

redColor='\033[0;31m'
greenColor='\033[0;32m'
yellowColor='\033[0;33m'
blueColor='\033[0;34m'
noColor='\033[0m'

errorCount=0
successCount=0
args=( "$@" )

sf() {
    echo -e "${blueColor}sf $@${noColor}"
    if [[ $args =~ "-d" ]]
    then
        command sf -debug "$@"
    else
        command sf "$@"
    fi    

    if [ $? != 0 ]
    then
        handleFailure
    else
        handleSuccess
    fi
}

handleFailure() {
    echo -e "${redColor}Fail${noColor}\n"
    let errorCount=errorCount+1
    if ! [[ $args =~ "-c" ]]
    then
        exit 1
    fi    
}

handleSuccess() {
    echo -e "${greenColor}Pass${noColor}\n"
    let successCount=successCount+1
}

report() {
    echo -e "${greenColor}Pass: ${successCount} ${redColor}Fail: ${errorCount}${noColor}"
    if [[ "$errorCount" -eq 0 ]]
    then
        echo -e "${greenColor}Tests passed${noColor}"
        exit 0
    else
        echo -e "${redColor}Tests failed${noColor}"
        exit 1
    fi
}

t() {
    echo -e "\n${yellowColor}$@${noColor}\n"   
}

t Nuking
sf nuke -confirm

t Adding and removing workspaces
sf a w W1
sf rm w W1
sf a w W1
sf cx W1
sf rm w

t Adding and removing directories
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf cx W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf rm w

t Adding and removing repositories
sf a w W1
sf a r -workspace W1 github.com/lachdoug/shapeframe-apps
sf rm r -workspace W1 github.com/lachdoug/shapeframe-apps
sf cx W1
sf a r github.com/lachdoug/shapeframe-apps
sf rm r github.com/lachdoug/shapeframe-apps
sf a r -https github.com/lachdoug/shapeframe-apps
sf pull github.com/lachdoug/shapeframe-apps
sf rm w

t Adding and removing frames
sf a w W1
sf cx W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf rm f docker-local
sf a f docker-local
sf cx docker-local
sf rm f
sf cx ..
sf a f -workspace W1 docker-local
sf rm f -workspace W1 docker-local
sf rm w W1

t Adding and removing shapes
sf a w W1
sf cx W1
sf a d /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f docker-local
sf cx docker-local
sf a s apache
sf rm s apache
sf cx ..
sf a s -frame docker-local apache
sf rm s -frame docker-local apache
sf cx ..
sf a s -workspace W1 -frame docker-local apache
sf rm s -workspace W1 -frame docker-local apache
sf rm w W1

t Listing
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf ls w
sf ls f -all
sf ls s -all
sf ls sr -all
sf ls fr -all
sf ls f -workspace W1 
sf ls s -workspace W1 
sf ls sr -workspace W1 
sf ls fr -workspace W1 
sf cx W1
sf ls f
sf ls s
sf ls sr
sf ls fr
sf rm w

t Reading
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf g w W1
sf g f -workspace W1 docker-local
sf g s -workspace W1 -frame docker-local apache
sf cx W1
sf g w
sf g f docker-local
sf g s -frame docker-local apache
sf cx docker-local
sf g f
sf g s apache
sf cx apache
sf g s
sf rm w

t Inspecting
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sg i W1
sf cx W1
sf i
sf rm w

t Contexting
sf cx
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf cx W1
sf cx
sf cx docker-local
sf cx
sf cx apache
sf cx
sf rm w

t Configuring
sf a w W1
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf cg f -workspace W1 -frame docker-local blue foo
sf cg s -workspace W1 -frame docker-local -shape apache blue foo cool
sf cx W1
sf cg f -frame docker-local blue foo
sf cg s -frame docker-local -shape apache blue foo cool
sf cx docker-local
sf cg s -shape apache blue foo cool
sf rm w

t Labelling
sf a w -about "Workspace One" W1 
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf ll w -name W1A -about "Workspace OneA" W1
sf ll f -workspace W1A -name docker-localA -about "Deploy containers on local Docker EngineA" docker-local
sf ll s -workspace W1A -frame docker-localA -name apacheA -about "Apache (Web Server)A" apache
sf ll w -name W1B -about "Workspace OneB" W1A
sf ll f -workspace W1B -name docker-localB -about "Deploy containers on local Docker EngineB" docker-localA
sf ll s -workspace W1B -frame docker-localB -name apacheB -about "Apache (Web Server)B" apacheA
sf cx W1B
sf ll w -name W1C -about "Workspace OneC"
sf ll f -name docker-localC -about "Deploy containers on local Docker EngineC" docker-localB
sf ll s -frame docker-localC -name apacheC -about "Apache (Web Server)C" apacheB
sf cx docker-localC
sf ll f -name docker-localD -about "Deploy containers on local Docker EngineD"
sf ll s -name apacheD -about "Apache (Web Server)D" apacheC
sf cx apacheD
sf ll s -name apacheE -about "Apache (Web Server)E"
sf rm w

t Orchestrating
sf a w -about "Workspace One" W1 
sf a d -workspace W1 /home/lachlan/Documents/play/go/sf-repos/shapeframe-apps
sf a f -workspace W1 docker-local
sf a s -workspace W1 -frame docker-local apache
sf o -workspace W1 -frame docker-local
sf cx W1
sf o -frame docker-local
sf cx docker-local
sf o

t Helping
sf -help
sf list -help
sf ls workspaces -help
sf ls frames -help
sf ls shapes -help
sf ls framers -help
sf ls shapers -help
sf label -help
sf ll shape -help
sf ll frame -help
sf ll workspace -help
sf configure -help
sf cg shape -help
sf cg frame -help
sf inspect -help
sf add -help
sf a shape -help
sf a frame -help
sf a workspace -help
sf a repository -help
sf a directory -help
sf remove -help
sf rm shape -help
sf rm frame -help
sf rm workspace -help
sf rm repository -help
sf rm directory -help
sf pull -help
sf context -help
sf orchestrate -help
sf nuke -help

report
