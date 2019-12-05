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

(defun find-noun-and-verb (ints i j)

    ;; Must keep track and recurse on copy of original ints list as
    ;; process-list alters the list itself
    (let ((orig-ints (copy-list ints)))

        (setf (nth 1 ints) i)
        (setf (nth 2 ints) j)
        (process-list ints 0)
        (if (eq (nth 0 ints) 19690720)
            (return-from find-noun-and-verb (list i j))
        )

        ;; Recurse or return nil if not found
        (cond
            ((and (eq i 99) (eq j 99)) nil)
            ((eq j 99) (find-noun-and-verb orig-ints (+ i 1) 0))
            (t (find-noun-and-verb orig-ints i (+ j 1)))
        )
    )
)

;; Part 1 Solution
(print (process-list '(1 12 2 3 1 1 2 3 1 3 4 3 1 5 0 3 2 1 9 19 1 10 19 23 2 9 23 27 1 6 27 31 2 31 9 35 1 5 35 39 1 10 39 43 1 10 43 47 2 13 47 51 1 10 51 55 2 55 10 59 1 9 59 63 2 6 63 67 1 5 67 71 1 71 5 75 1 5 75 79 2 79 13 83 1 83 5 87 2 6 87 91 1 5 91 95 1 95 9 99 1 99 6 103 1 103 13 107 1 107 5 111 2 111 13 115 1 115 6 119 1 6 119 123 2 123 13 127 1 10 127 131 1 131 2 135 1 135 5 0o 99 2 14 0 0) 0))

;; Part 2 Solution
(print (find-noun-and-verb '(1 0 0 3 1 1 2 3 1 3 4 3 1 5 0 3 2 1 9 19 1 10 19 23 2 9 23 27 1 6 27 31 2 31 9 35 1 5 35 39 1 10 39 43 1 10 43 47 2 13 47 51 1 10 51 55 2 55 10 59 1 9 59 63 2 6 63 67 1 5 67 71 1 71 5 75 1 5 75 79 2 79 13 83 1 83 5 87 2 6 87 91 1 5 91 95 1 95 9 99 1 99 6 103 1 103 13 107 1 107 5 111 2 111 13 115 1 115 6 119 1 6 119 123 2 123 13 127 1 10 127 131 1 131 2 135 1 135 5 0 99 2 14 0 0) 0 0))
