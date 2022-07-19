mod math;

use std::ffi::CStr;
use std::slice;

#[no_mangle]
pub extern "C" fn sum(n: *const i32, len: libc::size_t) -> i32 {
    let numbers = unsafe {
        assert!(!n.is_null());
        slice::from_raw_parts::<i32>(n, len as usize)
    };

    return math::vector_sum(numbers.to_vec()).unwrap();
}

#[no_mangle]
pub extern "C" fn hello(name: *const libc::c_char) {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("Hello {}!", str_name);
}

#[cfg(test)]
pub mod test {

    use crate::{hello, sum};
    use std::ffi::CString;

    #[test]
    fn simulated_main_function() {
        hello(CString::new("John Smith").unwrap().into_raw());
    }

    #[test]
    fn sum_test() {
        let arr: [i32; 8] = [1, 2, 3, 4, 5, 6, 7, 8];
        let r = sum(&arr as *const i32, 8);
        println!("{}", r)
    }
}
