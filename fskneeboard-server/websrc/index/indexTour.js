import "shepherd.js/dist/css/shepherd.css";
import Shepherd from "shepherd.js";

const tour = function() {
    const tour = new Shepherd.Tour({
        defaultStepOptions: {
          cancelIcon: {
            enabled: true
          },
          classes: 'class-1 class-2',
          scrollTo: { behavior: 'smooth', block: 'center' }
        }
      });
      
      tour.addStep({
        title: 'Creating a Shepherd Tour',
        text: `Creating a Shepherd tour is easy. too!\
        Just create a \`Tour\` instance, and add as many steps as you want.`,
        /*attachTo: {
          element: '#switch_map',
          on: 'bottom'
        },*/
        buttons: [
          {
            action() {
              return this.back();
            },
            classes: 'shepherd-button-secondary',
            text: 'Back'
          },
          {
            action() {
              return this.next();
            },
            text: 'Next'
          }
        ],
        id: 'creating'
      });
      
      tour.start();
}

export default tour;