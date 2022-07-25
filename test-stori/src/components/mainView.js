import ListNewsletter from './list/listNewsletter'
import AddNewsletter from './input/addNewsletter';
import ModalAddEmail from './modal/addEmail';

function MainView() {
  return (
    <>
    <ModalAddEmail />
    <AddNewsletter />
    <ListNewsletter />
    </>
  );
}

export default MainView;
