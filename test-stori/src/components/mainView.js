import ListNewsletter from './list/listNewsletter'
import AddNewsletter from './input/addNewsletter';
import ModalAddEmail from './modal/addEmail';

function MainView() {
  return (
    <>
    <AddNewsletter />
    <ListNewsletter />
    </>
  );
}

export default MainView;
