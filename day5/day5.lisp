;; This function is a modified version of the original process-list from day 2.
;; This function takes a list of integers as input and the index to start iterating
;; over (should be 0), and returns the altered array after all opcode operations are
;; parsed and applied to the list
(defun process-list (ints x)
    (if (and (not (eq (nth x ints) 99)) (<= x (list-length ints)))
        (progn
            ;; Addition opcode - store sum of next 2 positions in index = value of
            ;; element 3 positions forward
            (if (eq (nth x ints) 1)
                (setf (nth (nth (+ x 3) ints) ints) (+ (nth (nth (+ x 1) ints) ints) (nth (nth (+ x 2) ints) ints)))

                ;; Else if product opcode -store product of next 2 positions in index =
                ;; value of element 3 pos forward
                (if (eq (nth x ints) 2)
                    (setf (nth (nth (+ x 3) ints) ints) (* (nth (nth (+ x 1) ints) ints) (nth (nth (+ x 2) ints) ints)))
                )
            )

            ;; recurse on next opcode when operation is completed
            (process-list ints (+ x 4))
        )
        ;; When if statement fails, base case is reached, must return altered list
        (return-from process-list ints)
    )
)

;; Part 1 Solution
(let ((in (open "./puzzleinput.txt")))
    (let ((ints-string (read-line in)))
        (ql:quickload "cl-ppcre")
        (let ((ints (cl-ppcre:split "," ints-string)))
            (print ints)
        )
    )
    (close in)
)
